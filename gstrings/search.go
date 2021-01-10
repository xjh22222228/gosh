// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.

package gstrings

import (
	"regexp"
)

func Search(s, regex string) int {
	v, err := regexp.Compile(regex)

	if err != nil {
		return -1
	}

	result := v.FindStringIndex(s)

	if len(result) >= 1 {
		return result[0]
	}

	return -1
}
