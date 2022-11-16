package commands

import (
	"fmt"

	"github.com/cosmopool/artifex/internal/interfaces"
	"github.com/cosmopool/artifex/internal/terminal"
	"github.com/urfave/cli/v2"
)

func InitApp(deps *interfaces.Dependencies) *cli.App {
	return &cli.App{
		Name:  "artifex",
		Usage: "Create boilerplate files and folders for clean architecture.",
		Commands: []*cli.Command{
			InitCreateCommand(deps),
			InitConfigCommand(deps),
		},
	}
}

func readInput(defaultVal, option string, config interfaces.IConfiguration) error {
	var value string

	input := terminal.Read()
	if input == "" {
		value = defaultVal
	} else if defaultVal == "" {
		return fmt.Errorf("Default value for '%s' is empty", option)
	} else {
		value = input
	}

	config.SetOption(option, value)
	return nil
}
