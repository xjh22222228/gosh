// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// Since: v0.0.2

package gstr

import (
	"strings"
)

func Includes(s, searchString string, position ...int) bool {
	_position := 0

	if len(position) > 0 {
		_position = position[0]
	}

	if _position < 0 {
		_position = 0
	}

	if _position > len(s) - 1 {
		return false
	}

	s = s[_position:]

	return strings.Contains(s, searchString)
}
