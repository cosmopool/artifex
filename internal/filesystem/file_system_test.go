package filesystem

import (
	"testing"

	"github.com/cosmopool/artifex/internal/logger"
	"github.com/stretchr/testify/assert"
)

var fs = FS{
	log: logger.New(false),
}

func TestMakePathReturnPathWithNoBreakline(t *testing.T) {
	dir := "/test/dir\n"
	actual := fs.GeneratePath(dir)
	assert.Equal(t, "/test/dir", actual)
}

func TestMakePathRemoveAllMultipleBreakline(t *testing.T) {
	dir := "/test/\ndir\n"
	actual := fs.GeneratePath(dir)
	assert.Equal(t, "/test/dir", actual)
}

func TestMakePathRemoveDoubleSlashSeparator(t *testing.T) {
	dir := "/test//dir/\n"
	actual := fs.GeneratePath(dir)
	assert.Equal(t, "/test/dir/", actual)
}

func TestMakeFilePathReturnsCorrectWithNoPrefix(t *testing.T) {
	dir := "/test/dir/\n"
	template := "ConfigRepository"
	file := "file"
	extension := "extension"
	actual := fs.GenerateFilePath(dir, template, file, extension)
	assert.Equal(t, "/test/dir/file_repository.extension", actual)
}

func TestMakeFilePathReturnsCorrectWithPrefix(t *testing.T) {
	dir := "/test/dir/\n"
	template := "InterfaceConfigRepository"
	file := "file"
	extension := "extension"
	actual := fs.GenerateFilePath(dir, template, file, extension)
	assert.Equal(t, "/test/dir/interface_file_repository.extension", actual)
}
