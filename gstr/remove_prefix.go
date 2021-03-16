// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// https://docs.python.org/3/library/stdtypes.html#str.removeprefix
// Since: v0.0.5

package gstr

func RemovePrefix(str, prefix string) string {
	prefixLen := len(prefix)

	if prefixLen == 0 {
		return str
	}

	if str[0: prefixLen] == prefix {
		return str[prefixLen:]
	}

	return str
}
