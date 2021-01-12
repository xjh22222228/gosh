// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// Since: v0.0.3

package gstrings

func StartsWith(s, searchString string, position ...int) bool {
	_position := 0

	if len(position) > 0 {
		_position = position[0]
	}

	if _position < 0 {
		_position = 0
	}

	return s[_position: _position + len(searchString)] == searchString
}
