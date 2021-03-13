// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// Since: v0.0.2

package gstrings

import "math"

func Slice(s string, beginIndex int, endIndex ...int) string {
	strLen := len(s)
	endIdx := strLen

	if len(endIndex) > 0 {
		endIdx = endIndex[0]
	}

	if beginIndex < 0 {
		beginIndex = strLen + beginIndex

		if beginIndex < -strLen {
			beginIndex = int(math.Max(0, float64(beginIndex)))
		}
	}

	if beginIndex > strLen {
		beginIndex = strLen
	}

	if endIdx < 0 {
		endIdx = strLen + endIdx
	}

	if endIdx > strLen {
		endIdx = strLen
	}

	return s[beginIndex: endIdx]
}
