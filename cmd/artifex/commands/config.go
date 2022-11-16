package commands

import (
	"fmt"

	"github.com/cosmopool/artifex/internal/interfaces"
	"github.com/cosmopool/artifex/internal/terminal"
	"github.com/urfave/cli/v2"
)

func setFlags() []cli.Flag {
	usecase := &cli.BoolFlag{
		Name:    "usecase",
		Value:   false,
		Aliases: []string{"u"},
		Usage:   "Set 'usecase' configuration",
	}

	datasource := &cli.BoolFlag{
		Name:    "datasource",
		Value:   false,
		Aliases: []string{"d"},
		Usage:   "Set 'datasource' configuration",
	}

	repo := &cli.BoolFlag{
		Name:    "repository",
		Value:   false,
		Aliases: []string{"r"},
		Usage:   "Set 'repository' configuration",
	}

	init := &cli.BoolFlag{
		Name:    "init",
		Value:   false,
		Aliases: []string{"i"},
		Usage:   "Initialize all configuration. START HERE!",
	}

	clear := &cli.BoolFlag{
		Name:    "clear",
		Value:   false,
		Aliases: []string{"c"},
		Usage:   "Clear all configuration from repository",
	}

	return []cli.Flag{
		init,
		clear,
		usecase,
		repo,
		datasource,
	}
}

func InitConfigCommand(deps *interfaces.Dependencies) *cli.Command {
	return &cli.Command{
		Name:  "config",
		Flags: setFlags(),
		Usage: "Configure the layers path, interfaces and implementation names and more.",
		Action: func(ctx *cli.Context) error {
			startInit := ctx.Bool("init")
			shouldClear := ctx.Bool("clear")
			shouldConfigUsecase := ctx.Bool("usecase")
			shouldConfigRepository := ctx.Bool("repository")
			shouldConfigDatasource := ctx.Bool("datasource")

			if shouldClear {
				err := deps.Config.ClearAll()
				if err != nil {
					return err
				}
			}

			if shouldConfigUsecase || startInit {
				err := configUsecase(deps.Config)
				if err != nil {
					return err
				}
			}

			if shouldConfigRepository || startInit {
				err := configRepository(deps.Config)
				if err != nil {
					return err
				}
			}

			if shouldConfigDatasource || startInit {
				err := configDatasource(deps.Config)
				if err != nil {
					return err
				}
			}

			finishMessage()

			return nil
		},
	}
}

func finishMessage() {
	terminal.Clear()

	fmt.Println()
	fmt.Println("+――――――――――――――――――――――――――――――――――――――――――――――+")
	fmt.Println("│                                              │")
	fmt.Println("│            Configuration finished            │")
	fmt.Println("│          You can start create files!         │")
	fmt.Println("│                                              │")
	fmt.Println("+――――――――――――――――――――――――――――――――――――――――――――――+")
	fmt.Println("│    'artifex create --help' for more info     │")
	fmt.Println("+――――――――――――――――――――――――――――――――――――――――――――――+")
	fmt.Println()
}

func defaulValMessage() {
	fmt.Println()
	fmt.Println("+――――――――――――――――――――――――――――――――――――――――――――――+")
	fmt.Println("│                                              │")
	fmt.Println("│ Leave options EMPTY to use the DEFAULT value │")
	fmt.Println("│                                              │")
	fmt.Println("+――――――――――――――――――――――――――――――――――――――――――――――+")
	fmt.Println()
}
