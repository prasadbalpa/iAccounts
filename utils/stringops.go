package utils

import (
	"strings"
)

func FindAndReplace(targetstring string, find string, replace string) string {

	return (strings.Replace(targetstring, find, replace, -1))

}
