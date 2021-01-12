// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Array/fill
// Since: v0.0.3

package gslice

func Fill(elems []string, value string, start int, end ...int) []string {
	endIdx := len(elems)
	elLen := len(elems)

	if len(end) > 0 {
		endIdx = end[0]
	}

	if start < 0 {
		start = elLen + start
	}

	if endIdx < 0 {
		endIdx = elLen + endIdx
	}

	for i := 0; i < elLen; i++ {
		if i >= start && i <= endIdx {
			elems[i] = value
		}
	}

	return elems
}

func FillInt(elems []int, value int, start int, end ...int) []int {
	endIdx := len(elems)
	elLen := len(elems)

	if len(end) > 0 {
		endIdx = end[0]
	}

	if start < 0 {
		start = elLen + start
	}

	if endIdx < 0 {
		endIdx = elLen + endIdx
	}

	for i := 0; i < elLen; i++ {
		if i >= start && i <= endIdx {
			elems[i] = value
		}
	}

	return elems
}

func FillAny(elems []interface{}, value interface{}, start int,
end ...int) []interface{} {
	endIdx := len(elems)
	elLen := len(elems)

	if len(end) > 0 {
		endIdx = end[0]
	}

	if start < 0 {
		start = elLen + start
	}

	if endIdx < 0 {
		endIdx = elLen + endIdx
	}

	for i := 0; i < elLen; i++ {
		if i >= start && i <= endIdx {
			elems[i] = value
		}
	}

	return elems
}
