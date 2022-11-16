package commands

import (
	"fmt"

	"github.com/cosmopool/artifex/internal/configuration"
	"github.com/cosmopool/artifex/internal/interfaces"
	"github.com/cosmopool/artifex/internal/terminal"
)

func configRepository(config interfaces.IConfiguration) (err error) {
	var option, defaultVal string

	terminal.Clear()

	defaulValMessage()

	fmt.Println()
	fmt.Println()
	fmt.Println("     ━━━━━━━━━━")
	fmt.Println("━━━━ REPOSITORY ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("     ━━━━━━━━━━")
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("╭→ FILENAME ━━━")
	fmt.Println("│")
	fmt.Println("│ The filename template must contain 'Config' word.")
	fmt.Println("│")
	fmt.Println("│ When generating files, the 'Config' word will be")
	fmt.Println("│ replaced with given 'name' argument.")
	fmt.Println("│")
	fmt.Println("│ e.g.:")
	fmt.Println("│    template:   ConfigUsecase")
	fmt.Println("│     command:   artifex create -u -n delete")
	fmt.Println("│      output:   delete_usecase.dart")
	fmt.Println("│")
	fmt.Println("│")
	fmt.Println("│")
	fmt.Println("│")

	// REPOSITORY IMPLEMENTATION
	defaultVal = "/src/infra/repositories/"
	fmt.Printf("├ Implementation path (default: %s): ", defaultVal)
	option = configuration.REPOSITORY_IMPLEMENTATION_PATH
	err = readInput(defaultVal, option, config)
	if err != nil {
		return err
	}

	// REPOSITORY INTERFACE
	defaultVal = "/src/domain/repositories/"
	fmt.Printf("│\n╰ Interface path (default: %s): ", defaultVal)
	option = configuration.REPOSITORY_INTERFACE_PATH
	err = readInput(defaultVal, option, config)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("╭→ PATH")
	fmt.Println("│")
	fmt.Println("│ The path where to save created files for the layer.")
	fmt.Println("│")
	fmt.Println("│")
	fmt.Println("│")

	// REPOSITORY IMPLEMENTATION NAME TEMPLATE
	defaultVal = "ConfigRepositoryImpl"
	fmt.Printf("├ Implementation name. (default: '%s'): ", defaultVal)
	option = configuration.REPOSITORY_IMPLEMENTATION
	err = readInput(defaultVal, option, config)
	if err != nil {
		return err
	}

	// REPOSITORY INTERFACE NAME TEMPLATE
	defaultVal = "ConfigRepository"
	fmt.Printf("│\n╰ Interface name. (default: '%s'): ", defaultVal)
	option = configuration.REPOSITORY_INTERFACE
	err = readInput(defaultVal, option, config)
	if err != nil {
		return err
	}

	return
}
