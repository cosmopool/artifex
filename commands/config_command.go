package commands

import (
	"fmt"

	"github.com/cosmopool/artifex/core"
	"github.com/spf13/cobra"
)

var repositoryConfig *bool
var config = core.GetConfig()

func configInit() {
	repositoryConfig = configCmd.Flags().BoolP("repository", "u", false, "Configure 'repositorys' path")
	repositoryConfig = configCmd.Flags().BoolP("list", "l", false, "List all configuration values")
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
		boolean := true
		configRepository(&boolean)
		configRepositoryImplementation(&boolean)
		configRepositoryInterface(&boolean)
	},
}

func configRepository(shouldConfig *bool) {
	if *shouldConfig {
		fmt.Print("The implementation path for repository (e.g. /src/infra/repositories/): ")
		str := core.ReadTerminal()
		c := *config
		c.SetOption(core.REPOSITORY_IMPLEMENTATION_PATH, str)
		fmt.Println()
	}
}

func configRepositoryInterface(shouldConfig *bool) {
	if *shouldConfig {
		fmt.Println("Write a exemple of a REPOSITORY INTERFACE name")
		// fmt.Println("It must be a repository named 'Config' or 'config'")
		fmt.Print("(e.g. 'ConfigRepositoryInterface'): ")
		str := core.ReadTerminal()
		c := *config
		c.SetOption(core.REPOSITORY_INTERFACE, str)
		fmt.Println()
	}
}

func configRepositoryImplementation(shouldConfig *bool) {
	if *shouldConfig {
		fmt.Println("Write a exemple of a REPOSITORY IMPLEMENTATION name")
		// fmt.Println("It must be a repository named 'Config' or 'config'")
		fmt.Print("(e.g. 'ConfigRepositoryImplementation'): ")
		str := core.ReadTerminal()
		c := *config
		c.SetOption(core.REPOSITORY_IMPLEMENTATION, str)
		fmt.Println()
	}
}
