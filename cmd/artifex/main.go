package main

import (
	"os"

	"github.com/cosmopool/artifex/cmd/artifex/commands"
	"github.com/cosmopool/artifex/internal/configuration"
	"github.com/cosmopool/artifex/internal/filesystem"
	"github.com/cosmopool/artifex/internal/git"
	"github.com/cosmopool/artifex/internal/interfaces"
	"github.com/cosmopool/artifex/internal/logger"
)

func main() {
	deps := dependencies()
	log := deps.Log
	deps.Config.Init()

	app := commands.InitApp(deps)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	deps.Config.SaveAllOptions()
}

func dependencies() *interfaces.Dependencies {
	log := logger.New(false)

	gitCmd := git.New(log)

	options := make([]configuration.Option, 0, 30)
	conf := configuration.New(gitCmd, &options, log)

	fs := filesystem.New(log)

	return &interfaces.Dependencies{
		Config: conf,
		FS:     fs,
		Log:    log,
	}
}
