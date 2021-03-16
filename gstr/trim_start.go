// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// Since: v0.0.1

package gstr

import "regexp"

func TrimStart(str string) string {
	regex, _ := regexp.Compile(`^\s+`)

	return regex.ReplaceAllString(str, "")
}
