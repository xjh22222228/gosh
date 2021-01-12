// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// Since: v0.0.3

package grandom

func ShuffleInt(numbers []int) {
	length := len(numbers)
	var temp int

	for length > 0 {
		length -= 1
		random := RandInt(0, length)
		temp = numbers[length]
		numbers[length] = numbers[random]
		numbers[random] = temp
	}
}

func Shuffle(str []string) {
	length := len(str)
	var temp string

	for length > 0 {
		length -= 1
		random := RandInt(0, length)
		temp = str[length]
		str[length] = str[random]
		str[random] = temp
	}
}

func ShuffleAny(any []interface{}) {
	length := len(any)
	var temp interface{}

	for length > 0 {
		length -= 1
		random := RandInt(0, length)
		temp = any[length]
		any[length] = any[random]
		any[random] = temp
	}
}
