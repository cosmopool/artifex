package configuration

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
)

func New(gitCmd GitInterface, options *[]Option, log *zap.SugaredLogger) *Config {
	return &Config{
		GitCmd:  gitCmd,
		Options: options,
		log:     log,
	}
}

type Option struct {
	Name  string
	Value string
}

// GitExec is separate so we can inject in Git and test all methods
// This interface has only one method: execute git command and return it's error
type GitInterface interface {
	Do(...string) (string, error)
	GetAllOptions() ([]string, error)
	SetAllOptions([]string) error
	CheckIfIsValidRepo() (bool, error)
}

type Config struct {
	GitCmd  GitInterface
	Options *[]Option
	log     *zap.SugaredLogger
}

func (c *Config) Init() {
	isValid, err := (c.GitCmd).CheckIfIsValidRepo()
	if err != nil {
		c.log.Errorln("Could not check if current path is a valid git repository")
		c.log.Errorln()
		c.log.Fatalln("Use 'artifex config --help' to see more options")
	}
	if !isValid {
		c.log.Errorln("Current path is NOT a valid git repository")
		c.log.Errorln("Check if you are in the correct path or use 'git init'")
		c.log.Errorln()
		c.log.Fatalln("Use 'artifex config --help' to see more options")
	}

	c.parseOptions()
	//
	// emptyOptions := c.isRepoConfigured()
	// if len(emptyOptions) != 0 {
	// 	c.log.Errorln("[ERROR] Found empty option.")
	// 	c.log.Errorln("Set the following options with 'artifex config -s [option] [value]' " +
	// 		"and try again.")
	// 	for _, val := range emptyOptions {
	// 		c.log.Errorln(val.Name)
	// 	}
	// 	c.log.Errorln()
	// 	c.log.Fatalln("Use 'artifex config --help' to see more options")
	// }
}

// IsRepoConfigured check if every configuration has some value.
// Return a array of all empty config
func (c *Config) IsRepoConfigured() (emptyOptions []string) {
	emptyOptions = make([]string, 0, len(*c.Options))

	for _, config := range *c.Options {
		isConfigNotSet := config.Value == ""
		if isConfigNotSet {
			emptyOptions = append(emptyOptions, config.Name)
		}
	}

	return
}

// GetAllConfigValues parses `gitConfigs` array to an Config array
//
// This `gitConfigs` is the output of 'git config --list' command. Each item of
// `gitConfigs` is a line of the command output.
func (c *Config) parseOptions() {
	optionList := ConfigList()
	options := make([]Option, 0, len(optionList))
	gitOptions, _ := (c.GitCmd).GetAllOptions()

	for _, config := range gitOptions {
		isArtifexConfig := strings.HasPrefix(config, "artifex")

		if isArtifexConfig {
			name, value, found := strings.Cut(config, "=")
			if !found {
				c.log.Fatalf("Invalid artifex configuration: Missing '=' in config.\nGot: %s", config)
			}

			artifexConfig := Option{
				Name:  name,
				Value: value,
			}

			options = append(options, artifexConfig)
		}
	}

	*c.Options = options
}

// GetOption return the value of an option if exists, an error otherwise.
func (c *Config) GetOption(name string) (value string, err error) {
	for _, option := range *c.Options {
		if option.Name == name {
			value = option.Value
			return
		}
	}

	err = fmt.Errorf("No option found with name %s. Set the option first and try again", name)
	return
}

func (c *Config) updateOptionValue(name, value string) bool {
	for i, option := range *c.Options {
		if option.Name == name {
			o := *c.Options
			o[i] = Option{Name: name, Value: value}
			return true
		}
	}
	return false
}

// SetOption append an new Option to Config.Options
func (c *Config) SetOption(name, value string) {
	formattedValue := strings.Replace(value, "\n", "", -1)
	keyFoundAndUpdated := c.updateOptionValue(name, formattedValue)

	if !keyFoundAndUpdated {
		option := Option{Name: name, Value: formattedValue}
		*c.Options = append(*c.Options, option)
	}
}

// SaveAllOptions saves all options in local repo config file via GitCmd.Do.
func (c *Config) SaveAllOptions() (err error) {
	parsedOptions := []string{}
	for _, option := range *c.Options {
		optionParsed := option.Name + " " + option.Value
		parsedOptions = append(parsedOptions, optionParsed)
	}

	err = (c.GitCmd).SetAllOptions(parsedOptions)
	return
}

func (c *Config) ConfigList() []string {
	return []string{
		USECASE_IMPLEMENTATION_PATH,
		REPOSITORY_IMPLEMENTATION_PATH,
		DATASOURCE_IMPLEMENTATION_PATH,
		USECASE_IMPLEMENTATION,
		REPOSITORY_IMPLEMENTATION,
		DATASOURCE_IMPLEMENTATION,
		REPOSITORY_INTERFACE,
		USECASE_INTERFACE,
		DATASOURCE_INTERFACE,
		USECASE_INTERFACE_PATH,
		REPOSITORY_INTERFACE_PATH,
		DATASOURCE_INTERFACE_PATH,
	}
}

func (c *Config) ClearAll() (err error) {
	configList := c.ConfigList()
	for _, config := range configList {
		_, err = c.GitCmd.Do("config", "--unset", config)
		if err != nil {
			return
		}
	}
	return
}
