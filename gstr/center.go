// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// Since: v0.0.1

package gstr

import (
	"strings"
)

func Center(s string, width int, fillChar ...rune) string {
	if width <= len(s) {
		return s
	}

	fillC := []rune(" ")

	if len(fillChar) > 0 {
		fillC = []rune(string(fillChar[0]))
	}

	fillStr := string(fillC)

	var prefixStr = strings.Builder{}
	var suffixStr = strings.Builder{}

	for i := 1; i <= width - len(s); i += 1 {
		if i % 2 != 0 {
			prefixStr.WriteString(fillStr)
		} else {
			suffixStr.WriteString(fillStr)
		}
	}

	prefixStr.WriteString(s)
	prefixStr.WriteString(suffixStr.String())

	return prefixStr.String()
}
