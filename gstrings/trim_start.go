package gstrings

import "regexp"

func TrimStart(str string) string {
	regex, _ := regexp.Compile(`^\s+`)

	return regex.ReplaceAllString(str, "")
}
