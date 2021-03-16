// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/indexOf
// Since: v0.0.1

package gstr

import (
	"strings"
)

func IndexOf(str, searchStr string, fromIndex ...int) int {
	fromIdx := 0

	if len(fromIndex) > 0 {
		fromIdx = fromIndex[0]

		if fromIdx < 0 {
			fromIdx = 0
		}

		if fromIdx >= len(str) && searchStr != "" {
			return -1
		}
	}

	if searchStr == "" {
		if fromIdx == 0 || fromIdx < len(str) {
			return fromIdx
		}

		if fromIdx >= len(str) {
			return len(str)
		}
	}

	prefix := Slice(str, 0, fromIdx)
	suffix := Slice(str, fromIdx)

	resIdx := strings.Index(suffix, searchStr)

	if resIdx >= 0 {
		return len(prefix) + resIdx
	}

	return -1
}
