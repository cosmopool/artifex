package core

import "github.com/cosmopool/artifex/git"

var gitCmd = git.Git{}
var options = make([]Option, 0, 30)
var config = Config{GitCmd: &gitCmd, Options: &options}

func GetConfig() *Config {
	return &config
}
