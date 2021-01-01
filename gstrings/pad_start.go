package gstrings

import (
	"strings"
)

func PadStart(str string, targetLength int, padString ...string) string {
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

	return padStr[0: targetLength] + str
}
