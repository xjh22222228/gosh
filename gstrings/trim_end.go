package gstrings

import "regexp"

func TrimEnd(str string) string {
	regex, _ := regexp.Compile(`\s+$`)

	return regex.ReplaceAllString(str, "")
}
