package commands

import (
	"fmt"

	"github.com/cosmopool/artifex/internal/interfaces"
	"github.com/urfave/cli/v2"
)

func createFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:    "repository",
			Value:   false,
			Aliases: []string{"r"},
			Usage:   "create a repository",
		},
	}
}

func InitCreateCommand(deps *interfaces.Dependencies) *cli.Command {
	return &cli.Command{
		Name:  "create",
		Flags: createFlags(),
		Usage: "Create a single or multiple layer",
		Action: func(ctx *cli.Context) error {
			fmt.Print("creating layer: ", ctx.Args().First())
			if ctx.Bool("repository") {
				fmt.Println("repository")
			} else {
				fmt.Println("other")
			}
			return nil
		},
	}
}
