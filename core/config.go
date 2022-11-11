package core

import (
	"fmt"
	"log"
	"strings"

	"github.com/cosmopool/artifex/git"
)

type Option struct {
	Name  string
	Value string
}

type IConfiguration interface {
	Init()
	isRepoConfigured() []Option
	parseOptions() []Option
	SetOption(string)
	GetOption(string) (string, error)
	SaveAllOptions() error
}

type Config struct {
	GitCmd  git.IGit
	Options *[]Option
}

func (c *Config) Init() {
	isValid, err := c.GitCmd.CheckIfIsValidRepo()
	if err != nil {
		log.Println("Could not check if current path is a valid git repository")
		log.Println()
		log.Fatalln("Use 'artifex config --help' to see more options")
	}
	if !isValid {
		log.Println("Current path is NOT a valid git repository")
		log.Println("Check if you are in the correct path or use 'git init'")
		log.Println()
		log.Fatalln("Use 'artifex config --help' to see more options")
	}

	c.parseOptions()

	emptyOptions := c.isRepoConfigured()
	if len(emptyOptions) != 0 {
		log.Println("[ERROR] Found empty option.")
		log.Println("Set the following options with 'artifex config -s [option] [value]' " +
			"and try again.")
		for _, val := range emptyOptions {
			log.Println(val.Name)
		}
		log.Println()
		log.Fatalln("Use 'artifex config --help' to see more options")
	}
}

// IsRepoConfigured check if every configuration has some value.
// Return a array of all empty config
func (c *Config) isRepoConfigured() (emptyOptions []Option) {
	emptyOptions = make([]Option, 0, len(*c.Options))

	for _, config := range *c.Options {
		isConfigNotSet := config.Value == ""
		if isConfigNotSet {
			emptyOptions = append(emptyOptions, config)
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
	gitOptions, _ := c.GitCmd.GetAllOptions()

	for _, config := range gitOptions {
		isArtifexConfig := strings.HasPrefix(config, "artifex")

		if isArtifexConfig {
			name, value, found := strings.Cut(config, "=")
			if !found {
				err := fmt.Errorf("Invalid artifex configuration: Missing '=' in config.\nGot: %s", config)
				panic(err)
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

// SetOption append an new Option to Config.Options
func (c *Config) SetOption(name, value string) {
	formattedValue := strings.Replace(value, "\n", "", -1)
	option := Option{Name: name, Value: formattedValue}
	*c.Options = append(*c.Options, option)
}

// SaveAllOptions saves all options in local repo config file via GitCmd.Do.
func (c *Config) SaveAllOptions() (err error) {
	parsedOptions := []string{}
	for _, option := range *c.Options {
		optionParsed := option.Name + " " + option.Value
		parsedOptions = append(parsedOptions, optionParsed)
	}

	err = c.GitCmd.SetAllOptions(parsedOptions)
	return
}
