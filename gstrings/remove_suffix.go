// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// https://docs.python.org/3/library/stdtypes.html#str.removesuffix
// Since: v0.0.5

package gstrings

func RemoveSuffix(str, suffix string) string {
	suffixLen := len(suffix)

	if suffixLen == 0 {
		return str
	}

	if Slice(str, -suffixLen) == suffix {
		return Slice(str, 0, -suffixLen)
	}

	return str
}
