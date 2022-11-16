package commands

import (
	"fmt"

	"github.com/cosmopool/artifex/internal/configuration"
	"github.com/cosmopool/artifex/internal/interfaces"
)

func configPath(config interfaces.IConfiguration) (err error) {
	var option, defaultVal string

	fmt.Println()
	fmt.Println("╭→ USECASE Path")
	fmt.Println("│")

	// USECASE IMPLEMENTATION
	defaultVal = "/src/domain/usecases/"
	fmt.Printf("╰ Implementation path (default: %s): ", defaultVal)
	option = configuration.USECASE_IMPLEMENTATION_PATH
	err = readInput(defaultVal, option, config)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("╭→ REPOSITORY Path")
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
	fmt.Printf("╰ Interface path (default: %s): ", defaultVal)
	option = configuration.REPOSITORY_INTERFACE_PATH
	err = readInput(defaultVal, option, config)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("╭→ DATASOURCE Path")
	fmt.Println("│")

	// DATASOURCE IMPLEMENTATION
	defaultVal = "/src/external/datasources/"
	fmt.Printf("├ Implementation path (default: %s): ", defaultVal)
	option = configuration.DATASOURCE_IMPLEMENTATION_PATH
	err = readInput(defaultVal, option, config)
	if err != nil {
		return err
	}

	// DATASOURCE INTERFACE
	defaultVal = "/src/infra/datasources/"
	fmt.Printf("╰ Interface path (default: %s): ", defaultVal)
	option = configuration.DATASOURCE_INTERFACE_PATH
	err = readInput(defaultVal, option, config)
	if err != nil {
		return err
	}

	return
}
