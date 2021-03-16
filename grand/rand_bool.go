// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.

package grand

// Since: v0.0.3
func RandBool() bool {
	randI := Random() * float64(11)

	if randI > 5 {
		return true
	}

	return false
}
