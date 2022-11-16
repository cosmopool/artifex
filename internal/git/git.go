package git

import (
	"os/exec"
	"strings"

	"go.uber.org/zap"
)

func New(log *zap.SugaredLogger) *Git {
	return &Git{log: log}
}

type Git struct {
	log *zap.SugaredLogger
}

// Do executes a git command e.g.: 'git config --list'
func (g *Git) Do(command ...string) (output string, err error) {
	cmd := exec.Command("git", command...)
	g.log.Debugf("command: %s", strings.Join(cmd.Args, " "))
	stdoutBytes, err := cmd.Output()
	stdout := string(stdoutBytes)

	if stdout != "" {
		g.log.Debugf("output: %s", stdout)
	}
	if err != nil {
		g.log.Errorf("err: %s", cmd.Stderr)
	}

	return string(stdout), err
}

// GetAllOptions return all git repo options as a string list
func (g *Git) GetAllOptions() (allOptions []string, err error) {
	optionsOutput, err := g.Do("config", "--list")
	allOptions = strings.Split(optionsOutput, "\n")
	return
}

// SaveAllOptions saves all options in local repo config file via GitCmd.Do.
func (g *Git) SetAllOptions(options []string) (err error) {
	for _, option := range options {
		opt, val, _ := strings.Cut(option, " ")
		commands := []string{"config", opt, val}
		_, err = g.Do(commands...)
		if err != nil {
			return
		}
	}
	return
}

// CheckIfIsValidRepo
func (g *Git) CheckIfIsValidRepo() (isValid bool, err error) {
	output, err := g.Do("branch", "--list")
	isValid = !strings.Contains(output, "not a git repository")
	return
}
