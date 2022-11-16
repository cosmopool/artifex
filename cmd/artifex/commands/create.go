package commands

import (
	"fmt"

	"github.com/cosmopool/artifex/internal/configuration"
	"github.com/cosmopool/artifex/internal/interfaces"
	"github.com/urfave/cli/v2"
)

func createFlags() []cli.Flag {
	usecase := &cli.BoolFlag{
		Name:    "usecase",
		Value:   false,
		Aliases: []string{"u"},
		Usage:   "Create a 'usecase' layer",
	}

	datasource := &cli.BoolFlag{
		Name:    "datasource",
		Value:   false,
		Aliases: []string{"d"},
		Usage:   "Create a 'datasource' layer",
	}

	repo := &cli.BoolFlag{
		Name:    "repository",
		Value:   false,
		Aliases: []string{"r"},
		Usage:   "Create a 'repository' layer",
	}

	name := &cli.StringFlag{
		Name:    "name",
		Aliases: []string{"n"},
		Usage:   "Layer name",
	}

	return []cli.Flag{
		usecase,
		datasource,
		repo,
		name,
	}
}

func InitCreateCommand(deps *interfaces.Dependencies) *cli.Command {
	return &cli.Command{
		Name:  "create",
		Flags: createFlags(),
		Usage: "Create a single or multiple layer",
		Action: func(ctx *cli.Context) error {
			name := ctx.String("name")
			shouldClear := ctx.Bool("clear")
			shouldCreateUsecase := ctx.Bool("usecase")
			shouldCreateRepository := ctx.Bool("repository")
			shouldCreateDatasource := ctx.Bool("datasource")

			if shouldClear {
				err := deps.Config.ClearAll()
				if err != nil {
					return err
				}
			}

			if shouldCreateUsecase && name != "" {
				err := createUsecase(name, deps.Config, deps.FS)
				if err != nil {
					return err
				}
			}

			if shouldCreateRepository && name != "" {
				err := createRepository(name, deps.Config, deps.FS)
				if err != nil {
					return err
				}
			}

			if shouldCreateDatasource && name != "" {
				err := createDatasource(name, deps.Config, deps.FS)
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
}

func optionPathGeneration(
	fileName, implementationOption, template string,
	config interfaces.IConfiguration,
	fs interfaces.IFileSystem,
) (path string, err error) {
	dir, err := config.GetOption(implementationOption)
	if err != nil {
		return
	}

	path = fs.GeneratePath(dir)
	fileTemplate, err := config.GetOption(template)
	if err != nil {
		return
	}

	extension, err := config.GetOption(configuration.FILE_EXTENSION)
	if err != nil {
		return
	}

	path = fs.GenerateFilePath(path, fileTemplate, fileName, extension)
	return
}

func createUsecase(fileName string, config interfaces.IConfiguration, fs interfaces.IFileSystem) error {
	optionPath := configuration.USECASE_IMPLEMENTATION_PATH
	fileTemplate := configuration.USECASE_IMPLEMENTATION

	path, err := optionPathGeneration(fileName, optionPath, fileTemplate, config, fs)
	if err != nil {
		return err
	}
	fmt.Println(path)
	return nil
}

func createRepository(fileName string, config interfaces.IConfiguration, fs interfaces.IFileSystem) error {
	optionPath := configuration.REPOSITORY_IMPLEMENTATION_PATH
	fileTemplate := configuration.REPOSITORY_IMPLEMENTATION

	path, err := optionPathGeneration(fileName, optionPath, fileTemplate, config, fs)
	if err != nil {
		return err
	}
	fmt.Println(path)

	optionPath = configuration.REPOSITORY_IMPLEMENTATION_PATH
	fileTemplate = configuration.REPOSITORY_IMPLEMENTATION

	path, err = optionPathGeneration(fileName, optionPath, fileTemplate, config, fs)
	if err != nil {
		return err
	}
	fmt.Println(path)

	return nil
}

func createDatasource(fileName string, config interfaces.IConfiguration, fs interfaces.IFileSystem) error {
	optionPath := configuration.DATASOURCE_IMPLEMENTATION_PATH
	fileTemplate := configuration.DATASOURCE_IMPLEMENTATION

	path, err := optionPathGeneration(fileName, optionPath, fileTemplate, config, fs)
	if err != nil {
		return err
	}
	fmt.Println(path)

	optionPath = configuration.DATASOURCE_IMPLEMENTATION_PATH
	fileTemplate = configuration.DATASOURCE_IMPLEMENTATION

	path, err = optionPathGeneration(fileName, optionPath, fileTemplate, config, fs)
	if err != nil {
		return err
	}
	fmt.Println(path)

	return nil
}
