// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// Since: v0.0.5

package gfunctools

import (
	"github.com/xjh22222228/gosh/gstrings"
	"math"
	"strconv"
	"strings"
)

// Since: v0.0.5
func VersionCompare(v1, v2 string) int {
	v1 = gstrings.RemovePrefix(v1, "v")
	v1 = gstrings.RemovePrefix(v1, "V")
	v2 = gstrings.RemovePrefix(v2, "v")
	v2 = gstrings.RemovePrefix(v2, "V")

	version1 := strings.Split(v1, ".")
	version2 := strings.Split(v2, ".")
	length := int(math.Max(float64(len(version1)), float64(len(version2))))

	for len(version1) < length {
		version1 = append(version1, "0")
	}

	for len(version2) < length {
		version2 = append(version2, "0")
	}

	for i := 0; i < length; i++ {
		num1, _ := strconv.Atoi(version1[i])
		num2, _ := strconv.Atoi(version2[i])

		if num1 > num2 {
			return 1
		}

		if num1 < num2 {
			return -1
		}
	}

	return 0
}
