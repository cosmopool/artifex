package main

import (
	"os"

	"github.com/cosmopool/artifex/commands"
	"github.com/cosmopool/artifex/core"
	"github.com/cosmopool/artifex/git"
)

type exitCode int

const (
	exitOK     exitCode = 0
	exitError  exitCode = 1
	exitCancel exitCode = 2
	exitAuth   exitCode = 4
)

func main() {
	code := mainRun()
	os.Exit(int(code))
}

func mainRun() exitCode {
	git := git.Git{}
	options := make([]core.Option, 0, 30)
	config := core.Config{GitCmd: &git, Options: &options}
	config.Init()

	code := commands.Execute()
	return exitCode(code)
}
