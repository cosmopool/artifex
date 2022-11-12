package core

import (
	"fmt"
	"strings"

	"github.com/cosmopool/artifex/logger"
)

// GetTemplate returns a `configuration` template prefix and suffix.
//
// The `template` string must always contains 'Config' substring
// because it's used as the separator. It will return error if does
// not contains the substring.
func GetTemplate(template string) (prefix, suffix string, err error) {
	separator := "Config"
	prefix, suffix, found := strings.Cut(template, separator)
	if !found {
		err = fmt.Errorf("Not found '%s' word", separator)
	}
	return
}

func removeDoubleSlash(path string) string {
	temp := path
	for strings.Contains(temp, "//") {
		temp = strings.ReplaceAll(temp, "//", "/")
	}
	return temp
}

// sanitizePath add a trailing '/' if non is present to `path`
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
func MakePath(dir string) string {
	config := GetConfig()

	currentDir := config.CurrentDir
	if currentDir == "" {
		log := logger.GetLogger()
		log.Fatalln("Could not get current dir path.")
	}

	path := currentDir + sanitizePath(dir)
	path = strings.ReplaceAll(path, "\n", "")
	return removeDoubleSlash(path)
}

// MakePath add current dir path to `dir`
func MakeFilePath(dir, fileTemplate string) string {
	path := sanitizePath(dir) + sanitizePath(fileTemplate)
	path = strings.ReplaceAll(path, "\n", "")
	return removeDoubleSlash(path)
}

func ArgsContainsCommand(array []string, value string) bool {
	for _, element := range array {
		if element == value {
			return true
		}
	}
	return false
}
