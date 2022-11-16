package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/cosmopool/artifex/internal/configuration"
	"github.com/cosmopool/artifex/internal/interfaces"
)

func configDatasource(config interfaces.IConfiguration) (err error) {
	var option, defaultVal string

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("     ━━━━━━━━━━")
	fmt.Println("━━━━ DATASOURCE ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
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
	fmt.Printf("│\n╰ Interface path (default: %s): ", defaultVal)
	option = configuration.DATASOURCE_INTERFACE_PATH
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

	// DATASOURCE IMPLEMENTATION NAME TEMPLATE
	defaultVal = "ConfigDatasourceImpl"
	fmt.Printf("├ Implementation name. (default: '%s'): ", defaultVal)
	option = configuration.DATASOURCE_IMPLEMENTATION
	err = readInput(defaultVal, option, config)
	if err != nil {
		return err
	}

	// DATASOURCE INTERFACE NAME TEMPLATE
	defaultVal = "ConfigDatasource"
	fmt.Printf("│\n╰ Interface name. (default: '%s'): ", defaultVal)
	option = configuration.DATASOURCE_INTERFACE
	err = readInput(defaultVal, option, config)
	if err != nil {
		return err
	}

	return
}
