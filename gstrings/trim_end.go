// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.

package gstrings

import "regexp"

func TrimEnd(str string) string {
	regex, _ := regexp.Compile(`\s+$`)

	return regex.ReplaceAllString(str, "")
}
