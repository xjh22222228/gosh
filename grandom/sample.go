// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// https://www.w3schools.com/python/ref_random_sample.asp
// Since: v0.0.3

package grandom

import "log"

func Sample(elems []string, k int) []string {
	if k < 0 || k > len(elems) {
		log.Panicln("grandom.Sample: Sample larger than population or is" +
			" negative")
		log.Fatal()
	}
	copyElems := make([]string, len(elems))
	copy(copyElems, elems)

	list := make([]string, k)

	for i := 0; i < k; i++ {
		j := k - i - 1
		randInt := RandInt(0, j)

		temp := copyElems[randInt]
		list[i] = temp
		last := copyElems[j]
		copyElems[j] = temp
		copyElems[randInt] = last
	}

	return list
}

func SampleInt(elems []int, k int) []int {
	if k < 0 || k > len(elems) {
		log.Panicln("grandom.Sample: Sample larger than population or is" +
			" negative")
		log.Fatal()
	}
	copyElems := make([]int, len(elems))
	copy(copyElems, elems)

	list := make([]int, k)

	for i := 0; i < k; i++ {
		j := k - i - 1
		randInt := RandInt(0, j)

		temp := copyElems[randInt]
		list[i] = temp
		last := copyElems[j]
		copyElems[j] = temp
		copyElems[randInt] = last
	}

	return list
}

func SampleAny(elems []interface{}, k int) []interface{} {
	if k < 0 || k > len(elems) {
		log.Panicln("grandom.Sample: Sample larger than population or is" +
			" negative")
		log.Fatal()
	}
	copyElems := make([]interface{}, len(elems))
	copy(copyElems, elems)

	list := make([]interface{}, k)

	for i := 0; i < k; i++ {
		j := k - i - 1
		randInt := RandInt(0, j)

		temp := copyElems[randInt]
		list[i] = temp
		last := copyElems[j]
		copyElems[j] = temp
		copyElems[randInt] = last
	}

	return list
}
