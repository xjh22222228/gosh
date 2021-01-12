// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// Since: v0.0.2

package gstrings

func Slice(s string, beginIndex int, endIndex ...int) string {
	endIdx := len(s)

	if len(endIndex) > 0 {
		endIdx = endIndex[0]
	}

	if beginIndex < 0 {
		beginIndex = len(s) + beginIndex
	}

	if beginIndex > len(s) {
		beginIndex = len(s)
	}

	if endIdx < 0 {
		endIdx = len(s) + endIdx
	}

	if endIdx > len(s) {
		endIdx = len(s)
	}

	return s[beginIndex: endIdx]
}
