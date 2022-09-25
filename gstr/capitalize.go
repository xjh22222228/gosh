// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// Since: v0.0.5

package gstr

import (
	"strings"
)

func Capitalize[T string | []byte](str T) string {
	s := string(str)
	sLen := len(s)

	if sLen == 0 {
		return s
	}

	r := []rune(s)
	sFirst := strings.ToUpper(string(r[0]))
	return sFirst + strings.ToLower(string(r[1:]))
}
