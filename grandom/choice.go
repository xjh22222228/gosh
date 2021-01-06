package grandom

import "log"

// https://docs.python.org/zh-cn/3/library/random.html#random.choice
func Choice(str string) string {
	if str == "" {
		log.Panicln("Cannot choose from an empty sequence")
	}

	r := []rune(str)
	i := RandInt(0, len(r) - 1)

	return string(r[i])
}
