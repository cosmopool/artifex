package interfaces

import "go.uber.org/zap"

type IConfiguration interface {
	Init()
	SetOption(string, string)
	GetOption(string) (string, error)
	SaveAllOptions() error
	ConfigList() []string
	ClearAll() error
	IsRepoConfigured() []string
}

type IFileSystem interface {
	CreateDir(string) error
	CreateFile(path, fileName string) error
	GeneratePath(dir string) string
	GenerateFilePath(dir, fileTemplate, file, extension string) string
	TemplateSeparator(template string) (string, string)
}

type ILogger interface{}

type TerminalInterface interface {
	Read() string
}

type Dependencies struct {
	Config IConfiguration
	FS     IFileSystem
	Log    *zap.SugaredLogger
}
