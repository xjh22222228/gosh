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
