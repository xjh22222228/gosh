package gstrings

func Slice(s string, beginIndex int, endIndex ...int) string {
	_endIndex := len(s)

	if len(endIndex) > 0 {
		_endIndex = endIndex[0]
	}

	if beginIndex < 0 {
		beginIndex = len(s) + beginIndex
	}

	if beginIndex > len(s) {
		beginIndex = len(s)
	}

	if _endIndex < 0 {
		_endIndex = len(s) + _endIndex
	}

	if _endIndex > len(s) {
		_endIndex = len(s)
	}

	return s[beginIndex: _endIndex]
}
