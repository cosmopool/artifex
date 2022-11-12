package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeFilePathReturnPathWithNoBreakline(t *testing.T) {
	dir := "/test/dir\n"
	file := "/file.test"
	actual := MakeFilePath(dir, file)
	assert.Equal(t, "/test/dir/file.test", actual)
}

func TestMakeFilePathRemoveAllMultipleBreakline(t *testing.T) {
	dir := "/test/\ndir\n"
	file := "/file\n.test"
	actual := MakeFilePath(dir, file)
	assert.Equal(t, "/test/dir/file.test", actual)
}

func TestMakeFilePathRemoveDoubleSlashSeparator(t *testing.T) {
	dir := "/test/dir/\n"
	file := "/file.test"
	actual := MakeFilePath(dir, file)
	assert.Equal(t, "/test/dir/file.test", actual)
}

func TestMakeFilePathRemoveMultipleDoubleSlashSeparator(t *testing.T) {
	dir := "///test////dir/////\n"
	file := "file.test"
	actual := MakeFilePath(dir, file)
	assert.Equal(t, "/test/dir/file.test", actual)
}

func TestSanitize(t *testing.T) {
	dir := "///test///dir///\n"
	file := "//file.test"
	actual := MakeFilePath(dir, file)
	assert.Equal(t, "/test/dir/file.test", actual)
}
