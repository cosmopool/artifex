package commands

import (
	"fmt"

	"github.com/cosmopool/artifex/logger"
	"github.com/spf13/cobra"
)

func init() {
	configInit()
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.PersistentFlags().BoolVarP(&logger.DebugLevel, "verbose", "v", false, "Verbose output DEBUG level")
	rootCmd.PersistentFlags().BoolVarP(&logger.InfoLevel, "less-verbose", "n", false, "Less verbose output. Don't show INFO logs")
}

func Execute() int {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return 2
	}
	return 1
}

var rootCmd = &cobra.Command{
	Use:   "artifex",
	Short: "Creates the Clean Architecture boilerplate files for you.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		fmt.Println("Main func")
	},
}
