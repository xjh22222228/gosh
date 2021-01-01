package grandom

import (
	"math/rand"
	"time"
)

func Random() float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()
}
