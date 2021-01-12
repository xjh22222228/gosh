// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// Since: v0.0.1

package grandom

import (
	"math/rand"
	"time"
)

func Random() float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()
}
