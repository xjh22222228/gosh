package gstrings

import (
	"strings"
)

func Center(s string, width int, fillChar ...rune) string {
	if width <= len(s) {
		return s
	}

	_fillChar := []rune(" ")

	if len(fillChar) > 0 {
		_fillChar = []rune(string(fillChar[0]))
	}

	fillStr := string(_fillChar)

	var prefixStr = strings.Builder{}
	var suffixStr = strings.Builder{}

	for i := 1; i <= width - len(s); i += 1 {
		if i % 2 != 0 {
			prefixStr.WriteString(fillStr)
		} else {
			suffixStr.WriteString(fillStr)
		}
	}

	prefixStr.WriteString(s)
	prefixStr.WriteString(suffixStr.String())

	return prefixStr.String()
}
