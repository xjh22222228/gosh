// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// Since: v0.0.1

package gstrings

import "regexp"

// Since: v0.0.1
func TrimEnd(str string) string {
	regex, _ := regexp.Compile(`\s+$`)

	return regex.ReplaceAllString(str, "")
}
