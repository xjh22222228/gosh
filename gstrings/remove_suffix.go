package gstrings

// https://docs.python.org/3/library/stdtypes.html#str.removesuffix
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
