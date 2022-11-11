package git

import (
	"log"
	"os/exec"
	"strings"
)

// GitExec is separate so we can inject in Git and test all methods
// This interface has only one method: execute git command and return it's error
type IGit interface {
	Do(...string) (string, error)
	GetAllOptions() ([]string, error)
	SetAllOptions([]string) error
	CheckIfIsValidRepo() (bool, error)
}

type Git struct{}

// Do executes a git command e.g.: 'git config --list'
func (g *Git) Do(command ...string) (output string, err error) {
	cmd := exec.Command("git", command...)
	log.Printf("[DEBUG] gitDo command: %s", cmd.Args)
	stdoutBytes, err := cmd.Output()
	stdout := string(stdoutBytes)

	if stdout != "" {
		log.Printf("[DEBUG] gitDo output: %s", stdout)
	}
	if err != nil {
		log.Printf("[DEBUG] gitDo err: %s", err)
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
		_, err = g.Do("config", option)
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
