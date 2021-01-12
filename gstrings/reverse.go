// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// Since: v0.0.1

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
