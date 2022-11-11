package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of artifex",
	Long:  `All software has versions. This is artifex's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Kaio Delphino's artifex CLI software v0.0.1 -- HEAD")
	},
}
