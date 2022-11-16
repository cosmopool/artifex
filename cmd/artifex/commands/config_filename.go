package commands

import (
	"fmt"

	"github.com/cosmopool/artifex/internal/configuration"
	"github.com/cosmopool/artifex/internal/interfaces"
)

func configNameTemplate(config interfaces.IConfiguration) (err error) {
	var option, defaultVal string

	fmt.Println()
	fmt.Println()
	fmt.Println("╭→ DATASOURCE Filename")
	fmt.Println("│")

	// REPOSITORY IMPLEMENTATION NAME TEMPLATE
	defaultVal = "ConfigUsecase"
	fmt.Printf("Implementation name. Must contain 'Config'  (default: '%s'): ", defaultVal)
	fmt.Println()
	option = configuration.REPOSITORY_IMPLEMENTATION
	err = readInput(defaultVal, option, config)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println()
	fmt.Println("╭→ REPOSITORY Filename")
	fmt.Println("│")

	// REPOSITORY IMPLEMENTATION NAME TEMPLATE
	defaultVal = "ConfigRepositoryImplementation"
	fmt.Printf("Implementation name. Must contain 'Config'  (default: '%s'): ", defaultVal)
	fmt.Println()
	option = configuration.REPOSITORY_IMPLEMENTATION
	err = readInput(defaultVal, option, config)
	if err != nil {
		return err
	}

	// REPOSITORY INTERFACE NAME TEMPLATE
	defaultVal = "ConfigRepositoryInterface"
	fmt.Printf("Interface name. Must contain 'Config' (default: '%s'): ", defaultVal)
	fmt.Println()
	option = configuration.REPOSITORY_INTERFACE
	err = readInput(defaultVal, option, config)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println()
	fmt.Println("╭→ DATASOURCE Filename")
	fmt.Println("│")

	// DATASOURCE IMPLEMENTATION NAME TEMPLATE
	defaultVal = "ConfigDatasourceImplementation"
	fmt.Printf("Implementation name. Must contain 'Config'  (default: '%s'): ", defaultVal)
	fmt.Println()
	option = configuration.DATASOURCE_IMPLEMENTATION
	err = readInput(defaultVal, option, config)
	if err != nil {
		return err
	}

	// DATASOURCE INTERFACE NAME TEMPLATE
	defaultVal = "ConfigDatasourceInterface"
	fmt.Printf("Interface name. Must contain 'Config' (default: '%s'): ", defaultVal)
	fmt.Println()
	option = configuration.DATASOURCE_INTERFACE
	err = readInput(defaultVal, option, config)
	if err != nil {
		return err
	}

	return
}
