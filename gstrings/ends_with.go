// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.

package gstrings

func EndsWith(str, searchString string, length ...int) bool {
	_length := len(str)

	if len(length) > 0 {
		_length = length[0]
	}

	if _length > len(str) {
		_length = len(str)
	}

	left := _length - len(searchString)

	if left < 0 {
		left = 0
	}

	return str[left: _length] == searchString
}
