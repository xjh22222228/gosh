// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// Since: v0.0.5

package grand

import "math"

func Range(start, stop int, step ...int) []int {
	stepN := 1
	if len(step) > 0 {
		stepN = step[0]
	}

	if start > stop || stepN < 1 {
		return make([]int, 0)
	}

	total := int(math.Ceil(float64(stop - start) / float64(stepN)))
	nums := make([]int, total)

	for i := start; i < total; i += 1 {
		nums[i] = i * stepN
	}

	return nums
}
