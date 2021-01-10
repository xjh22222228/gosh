package gstrings

// https://docs.python.org/3/library/stdtypes.html#str.removeprefix
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
