// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// Since: v0.0.6

package gmath

import (
	"fmt"
	"math"
	"strconv"
)

func Round(x float64, precision ...int) float64 {
	pre := 0

	if len(precision) > 0 {
		pre = precision[0]

		if pre >= 0 {
			pre = int(math.Min(float64(pre), 292))
		} else {
			pre = 0
		}
	}

	format := "%." + strconv.Itoa(pre) + "f"

	v, err := strconv.ParseFloat(fmt.Sprintf(format, x), 64)

	if err != nil {
		return 0
	}

	return v
}
