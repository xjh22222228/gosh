// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.

package gmath

func Sum(num ...int) int {
	n := 0
	i := len(num)

	for i > 0 {
		i--
		n += num[i]
	}

	return n
}
