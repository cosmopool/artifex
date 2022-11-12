package commands

import (
	"fmt"

	"github.com/cosmopool/artifex/core"
	"github.com/spf13/cobra"
)

func configCreate() {
	// createAll = configCmd.Flags().BoolP("repository", "r", false, "Configure 'repositories' path")
}

var createCmd = &cobra.Command{
	Use:       "create",
	Short:     "Create files in layers",
	ValidArgs: []string{"all", "repository", "usecase", "datasource"},
	Args:      cobra.OnlyValidArgs,
	Long: "Create the interfaces and implementation files with name template" +
		"From configuration. Can create for single or multiple layers.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println(args)
		if core.ArgsContainsCommand(args, "all") {
			createAll()
		}
		if core.ArgsContainsCommand(args, "repository") {
			createRepository()
		}
		if core.ArgsContainsCommand(args, "usecase") {
			createUsecase()
		}
		if core.ArgsContainsCommand(args, "datasource") {
			createDatasource()
		}
	},
}

func createAll() {}

func createRepository() {
	config := core.GetConfig()
	// manager := dir.GetFileManager()

	dir, _ := config.GetOption(core.REPOSITORY_IMPLEMENTATION_PATH)
	fmt.Println(dir)
	path := core.MakePath(dir)
	path = core.MakeFilePath(path, "test.dart")
	fmt.Println(path)

	// cmd := exec.Command("pwd")
	// stdoutBytes, _ := cmd.Output()
	// currentDir := string(stdoutBytes)

	// path := strings.Replace(currentDir+dir, "\n", "", 2)
	// manager.CreateDir(path)
}

func createUsecase() {}

func createDatasource() {}
