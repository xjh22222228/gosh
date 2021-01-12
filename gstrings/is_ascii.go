// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// Since: v0.0.3

package gstrings

const (
	minASCII = '\u0000'
	maxASCII = '\u007F'
)

// https://docs.python.org/3/library/stdtypes.html#str.isascii
func IsAscii(str string) bool {
	r := []rune(str)
	i := 0
	rLen := len(r)

	for i < rLen {
		u32 := uint32(r[i])
		if u32 < minASCII || u32 > maxASCII {
			return false
		}

		i += 1
	}

	return true
}
