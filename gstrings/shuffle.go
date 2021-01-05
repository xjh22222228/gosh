package gstrings

import (
	"github.com/xjh22222228/gosh/grandom"
)

func Shuffle(str string) string {
	r := []rune(str)
	length := len(r)
	var temp int32

	for length > 0 {
		length -= 1
		random := grandom.RandInt(0, length)
		temp = r[length]
		r[length] = r[random]
		r[random] = temp
	}

	return string(r)
}
