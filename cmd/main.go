package main

import (
	"log"
	"os"

	"github.com/cosmopool/artifex/commands"
	"github.com/cosmopool/artifex/core"
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
	config := core.GetConfig()
	config.Init()

	code := commands.Execute()
	err := config.SaveAllOptions()
	if err != nil {
		log.Fatalln(err)
	}
	return exitCode(code)
}
