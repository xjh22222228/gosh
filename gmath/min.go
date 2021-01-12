// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// Since: v0.0.5

package gmath

func Min(n ...float64) float64 {
	var f float64
	i := len(n)

	if i > 0 {
		f = n[0]
	}

	for i > 0 {
		i--
		if n[i] < f {
			f = n[i]
		}
	}

	return f
}
