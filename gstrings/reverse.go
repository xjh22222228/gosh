package gstrings

func Reverse(str string) string {
	r := []rune(str)
	ns := ""

	for idx := range r {
		ns += string(r[len(r) - idx - 1])
	}

	return ns
}
