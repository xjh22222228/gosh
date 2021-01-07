package gstrings

import "strings"

func Reverse(str string) string {
	r := []rune(str)
	var ns strings.Builder

	for idx := range r {
		ns.WriteRune(r[len(r) - idx - 1])
	}

	return ns.String()
}
