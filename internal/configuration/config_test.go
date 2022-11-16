package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockGit struct {
	mock.Mock
}

func (m *MockGit) Do(command ...string) (string, error) {
	return m.Called().String(), nil
}

func (m *MockGit) GetAllOptions() ([]string, error) {
	args := m.Called()
	var checkInterface interface{} = args[0]
	allOptions := checkInterface.([]string)
	return allOptions, nil
}

func (m *MockGit) SetAllOptions(options []string) (err error) {
	m.Called(options)
	return nil
}

func (m *MockGit) CheckIfIsValidRepo() (bool, error) {
	return true, nil
}

// test parseOptions
func TestAllConfigValuesCanParseAllItems(t *testing.T) {
	gitOptions := []string{
		"artifex.path=/src",
		"artifex.domain=/src/domain",
		"artifex.external=/src/external",
	}

	gitCmd := new(MockGit)
	gitCmd.On("GetAllOptions").Return(gitOptions)

	options := make([]Option, 0, 10)
	config := Config{
		GitCmd:  gitCmd,
		Options: &options,
	}
	config.parseOptions()

	assert.Equal(t, len(*config.Options), len(gitOptions), "Could not parse config name correctly")
}

// test parseOptions
func TestAllConfigValuesCanParseConfigurationNameCorrectly(t *testing.T) {
	expectedName := "artifex.path"
	expectedValue := "/src"
	gitOptions := []string{
		expectedName + "=" + expectedValue,
	}

	gitCmd := new(MockGit)
	gitCmd.On("GetAllOptions").Return(gitOptions)

	options := make([]Option, 0, 10)
	config := Config{
		GitCmd:  gitCmd,
		Options: &options,
	}
	config.parseOptions()

	optionVal := *config.Options
	actual := optionVal[0].Name
	assert.Equal(t, expectedName, actual, "Could not parse config name correctly")
}

// test parseOptions
func TestAllConfigValuesCanParseConfigurationValueCorrectly(t *testing.T) {
	expectedName := "artifex.path"
	expectedValue := "/src"
	gitOptions := []string{
		expectedName + "=" + expectedValue,
	}

	gitCmd := new(MockGit)
	gitCmd.On("GetAllOptions").Return(gitOptions)

	options := make([]Option, 0, 10)
	config := Config{
		GitCmd:  gitCmd,
		Options: &options,
	}
	config.parseOptions()

	optionVal := *config.Options
	actual := optionVal[0].Value
	assert.Equal(t, expectedValue, actual, "Could not parse config name correctly")
}

// test parseOptions
func TestAllConfigValuesCanParseEmptyValue(t *testing.T) {
	gitOptions := []string{
		"artifex.path=",
	}

	gitCmd := new(MockGit)
	gitCmd.On("GetAllOptions").Return(gitOptions)

	options := make([]Option, 0, 10)
	config := Config{
		GitCmd:  gitCmd,
		Options: &options,
	}
	config.parseOptions()

	optionVal := *config.Options
	actual := optionVal[0].Value
	assert.Equal(t, "", actual, "Could not parse config name correctly")
}

// test GetOption
func TestTryGetNonExistentOptionReturnsError(t *testing.T) {
	options := make([]Option, 0, 10)
	config := Config{
		GitCmd:  new(MockGit),
		Options: &options,
	}

	_, err := config.GetOption("non.extistent.option")

	expectedErrorMsg := "No option found with name non.extistent.option. Set the option first and try again"
	assert.EqualErrorf(t, err, expectedErrorMsg, "Errors message string do not match")
}

// test GetOption
func TestGetOptionReturnsRightValue(t *testing.T) {
	name := "artifex.config.some"
	expected := "yes"

	opt := Option{Name: name, Value: expected}
	options := []Option{opt}

	config := Config{
		GitCmd:  new(MockGit),
		Options: &options,
	}
	actual, _ := config.GetOption(name)

	assert.Equal(t, expected, actual)
}

// test SetOption
func TestSetOptionAppendValue(t *testing.T) {
	name := "artifex.config.some"
	value := "yes"

	opt := Option{Name: "name", Value: "val"}
	options := []Option{opt}

	config := Config{
		GitCmd:  new(MockGit),
		Options: &options,
	}
	config.SetOption(name, value)

	assert.Equal(t, len(*config.Options), 2)
}

// test SetOption
func TestSetOptionShouldOverwriteOldOptionInsteadOfAppend(t *testing.T) {
	name := "artifex.config.some"
	value := "newVal"

	opt := Option{Name: name, Value: "oldVal"}
	options := []Option{opt}

	config := Config{
		GitCmd:  new(MockGit),
		Options: &options,
	}
	config.SetOption(name, value)

	assert.Equal(t, 1, len(*config.Options))
}

// test SetOption
func TestSetOptionShouldOverwriteOldValueInsteadOfAppend(t *testing.T) {
	name := "artifex.config.some"
	newVal := "newVal"

	opt := Option{Name: name, Value: "oldVal"}
	options := []Option{opt}

	config := Config{
		GitCmd:  new(MockGit),
		Options: &options,
	}
	config.SetOption(name, newVal)
	o := *config.Options
	option := o[0]

	assert.Equal(t, newVal, option.Value)
}

// test SaveAllOptions
func TestSaveAllOptionsSendCorrectParsedOptions(t *testing.T) {
	name := "artifex.config.some"
	value := "yes"
	args := []string{name + " " + value}

	opt := Option{Name: name, Value: value}
	options := []Option{opt}

	gitCmd := new(MockGit)
	gitCmd.On("SetAllOptions", args).Return(nil)

	config := Config{
		GitCmd:  gitCmd,
		Options: &options,
	}

	config.SaveAllOptions()
	actual := gitCmd.AssertCalled(t, "SetAllOptions", args)

	assert.Equal(t, true, actual)
}
