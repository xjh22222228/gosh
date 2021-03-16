// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// Since: v0.0.3

package gstr

import (
	"github.com/xjh22222228/gosh/grand"
)

func Shuffle(str string) string {
	r := []rune(str)
	length := len(r)
	var temp int32

	for length > 0 {
		length -= 1
		random := grand.RandInt(0, length)
		temp = r[length]
		r[length] = r[random]
		r[random] = temp
	}

	return string(r)
}
