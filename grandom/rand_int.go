package grandom

import (
	"math"
)

func RandInt(a, b int) int {
	randI := math.Floor(Random() * float64(b))

	for randI < float64(a) {
		return b
	}

	return int(randI)
}
