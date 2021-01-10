// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.

package grandom

import (
	"log"
	"math/rand"
	"time"
)

func RandInt(min, max int) int {
	if min == max {
		return min
	}

	if min > max {
		log.Panicf("min(%v) cannot be greater than max(%v)", min, max)
	}

	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max + 1 - min) + min
}
