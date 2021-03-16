// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// Since: v0.0.2

package gstr

import "strings"

func PadEnd(str string, targetLength int, padString ...string) string {
	padStr := " "

	if len(padString) > 0 {
		padStr = padString[0]
	}

	if len(str) > targetLength {
		return str
	}

	targetLength = targetLength - len(str)
	if targetLength > len(padStr) {
		padStr += strings.Repeat(padStr, targetLength)
	}

	return str + padStr[0: targetLength]
}
