// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Array/slice
// Since: v0.0.1

package gslice

import "math"

func Slice[T any](s []T, beginIndex int, endIndex ...int) []T {
	sLen := len(s)
	endIdx := sLen

	if len(endIndex) > 0 {
		endIdx = endIndex[0]
	}

	if beginIndex < 0 {
		beginIndex = sLen + beginIndex

		if beginIndex < -sLen {
			beginIndex = int(math.Max(0, float64(beginIndex)))
		}
	}

	if beginIndex > sLen {
		beginIndex = sLen
	}

	if endIdx < 0 {
		endIdx = sLen + endIdx
	}

	if endIdx > sLen {
		endIdx = sLen
	}

	copyElems := make([]T, endIdx-beginIndex)
	copy(copyElems, s[beginIndex:endIdx])
	return copyElems
}
