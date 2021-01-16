// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// Since: v0.0.3

package grandom

import (
	"math/rand"
	"time"
)

func RandInt(min, max int) int {
	if min == max {
		return min
	}

	if min > max {
		return 0
	}

	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max + 1 - min) + min
}
