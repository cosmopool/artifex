package filesystem

import (
	"os"
	"os/exec"
	"strings"

	"go.uber.org/zap"
)

func New(log *zap.SugaredLogger) *FS {
	return &FS{log: log}
}

type FS struct {
	log *zap.SugaredLogger
}

func (f *FS) CreateDir(dir string) error {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		f.log.Errorln("Encounter error trying to create dir:", dir)
		f.log.Fatalln("Error:", err)
	}
	return nil
}

func (f *FS) CreateFile(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		f.log.Errorln("Encounter error trying to create file:", fileName)
		f.log.Fatalln("Error:", err)
	}

	file.Close()
	return nil
}

func removeDoubleSlash(path string) string {
	temp := path
	for strings.Contains(temp, "//") {
		temp = strings.ReplaceAll(temp, "//", "/")
	}
	return temp
}

// sanitizePath add a trailing '/' if non is present in `path`
func sanitizePath(path string) string {
	isValidPath := strings.HasPrefix(path, "/")
	if isValidPath {
		return path
	} else {
		path = "/" + path
		return path
	}
}

// MakePath add current dir path to `dir`
func (f *FS) GeneratePath(dir string) string {
	currentDir, err := f.currentDir()
	if err != nil || currentDir == "" {
		f.log.Fatalln("Could not get current dir path.", err)
	}

	path := currentDir + sanitizePath(dir)
	path = strings.ReplaceAll(path, "\n", "")
	return removeDoubleSlash(path)
}

// MakeFilePath add current dir path and `fileTemplate` to `dir`
func (f *FS) GenerateFilePath(dir, fileTemplate, file, extension string) string {
	prefix, suffix := f.TemplateSeparator(fileTemplate)
	fileName := prefix + file + suffix + "." + extension
	path := sanitizePath(dir) + sanitizePath(fileName)
	path = strings.ReplaceAll(path, "\n", "")
	path = strings.ToLower(path)
	return removeDoubleSlash(path)
}

// TemplateSeparator returns a `configuration` template prefix and suffix.
//
// The `template` string must always contains 'Config' substring
// because it's used as the separator. It will return error if does
// not contains the substring.
func (f *FS) TemplateSeparator(template string) (prefix, suffix string) {
	separator := "Config"
	prefix, suffix, found := strings.Cut(template, separator)
	if !found {
		f.log.Fatalf("Not found '%s' word in filename template '%s'", separator, template)
	}

	if prefix != "" {
		prefix = prefix + "_"
	}
	if suffix != "" {
		suffix = "_" + suffix
	}

	return
}

func (*FS) currentDir() (currentDir string, err error) {
	cmd := exec.Command("pwd")
	stdoutBytes, err := cmd.Output()
	currentDir = string(stdoutBytes)
	return
}
