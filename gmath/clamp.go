// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// Since: v0.0.7

package gmath

func Clamp(number, lower, upper int) int {
    if number > upper {
        number = upper
    }

    if number < lower {
        number = lower
    }

    return number
}

func ClampFloat(number, lower, upper float64) float64 {
    if number > upper {
        number = upper
    }

    if number < lower {
        number = lower
    }

    return number
}