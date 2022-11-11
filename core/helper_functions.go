package core

import (
	"fmt"
	"strings"
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
