// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// Since: v0.0.5

package gmath

import "math"

func InRange(number, start float64, end ...float64) bool {
	var endNum float64

	if len(end) > 0 {
		endNum = end[0]
	} else {
		endNum = start
		start = 0
	}

	return number >= math.Min(start, endNum) && number < math.Max(start, endNum)
}
