package commands

import (
	"github.com/cosmopool/artifex/core"
	"github.com/spf13/cobra"
)

var repositoryConfig *bool
var shouldListConfig *bool
var config = core.GetConfig()

func configInit() {
	repositoryConfig = configCmd.Flags().BoolP("repository", "r", false, "Configure 'repositories' path")
	shouldListConfig = configCmd.Flags().BoolP("list", "l", false, "List all configuration values")
}

func existsElement(array []string, str string) bool {
	for _, val := range array {
		if val == str {
			return true
		}
	}
	return false
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the application",
	Long: `Configure the layers path, interfaces and implementation names and 
	more. All the configuration the application needs to run`,
	Run: func(cmd *cobra.Command, args []string) {
		configRepository(cmd)
		configRepositoryImplementation(cmd)
		configRepositoryInterface(cmd)
		listConfig(cmd)
	},
}

func listConfig(cmd *cobra.Command) {
	if *shouldListConfig {
		options := *config.Options
		for _, option := range options {
			cmd.Println(option.Name, option.Value)
		}
	}
}

func configRepository(cmd *cobra.Command) {
	if *repositoryConfig {
		cmd.Print("The implementation path for repository (e.g. /src/infra/repositories/): ")
		str := core.ReadTerminal()
		c := *config
		c.SetOption(core.REPOSITORY_IMPLEMENTATION_PATH, str)
		cmd.Println()
	}
}

func configRepositoryInterface(cmd *cobra.Command) {
	if *repositoryConfig {
		cmd.Println("Write a exemple of a REPOSITORY INTERFACE name")
		// cmd.Println("It must be a repository named 'Config' or 'config'")
		cmd.Print("(e.g. 'ConfigRepositoryInterface'): ")
		str := core.ReadTerminal()
		c := *config
		c.SetOption(core.REPOSITORY_INTERFACE, str)
		cmd.Println()
	}
}

func configRepositoryImplementation(cmd *cobra.Command) {
	if *repositoryConfig {
		cmd.Println("Write a exemple of a REPOSITORY IMPLEMENTATION name")
		// cmd.Println("It must be a repository named 'Config' or 'config'")
		cmd.Print("(e.g. 'ConfigRepositoryImplementation'): ")
		str := core.ReadTerminal()
		c := *config
		c.SetOption(core.REPOSITORY_IMPLEMENTATION, str)
		cmd.Println()
	}
}
