// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// Since: v0.0.5

package gstr

import "strings"

func UpperFirst(str string) string {
	sLen := len(str)
	if sLen == 0 {
		return str
	}

	sFirst := strings.ToUpper(string(str[0]))


	return sFirst + str[1:]
}
