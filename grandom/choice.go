// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.

package grandom

import "log"

// https://docs.python.org/zh-cn/3/library/random.html#random.choice
func Choice(str string) string {
	if str == "" {
		return ""
	}

	r := []rune(str)
	i := RandInt(0, len(r) - 1)

	return string(r[i])
}

func ChoiceInt(nums []int) int {
	if len(nums) == 0 {
		log.Panicln("Cannot choose from an empty sequence")
	}

	i := RandInt(0, len(nums) - 1)
	return nums[i]
}

func ChoiceStr(s []string) string {
	if len(s) == 0 {
		log.Panicln("Cannot choose from an empty sequence")
	}

	i := RandInt(0, len(s) - 1)
	return s[i]
}
