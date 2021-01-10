// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.

package gstrings

import (
	"strings"
)

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/lastIndexOf
func LastIndexOf(str, searchStr string, fromIndex ...int) int {
	fromIdx := len(str)

	if len(fromIndex ) > 0 {
		fromIdx = fromIndex[0]

		if fromIdx < 0 {
			fromIdx = 0
		}
	}

	if searchStr == "" {
		return fromIdx
	}

	newStr := Slice(str, 0, fromIdx + 1)

	return strings.LastIndex(newStr, searchStr)
}
