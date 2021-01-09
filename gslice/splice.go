package gslice

import (
	"math"
)

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Array/splice
func Splice(elems *[]string, start, deleteCount int, item ...string) []string {
	pElems := *elems
	elLen := len(pElems)
	itemLen := len(item)

	if start > elLen {
		start = elLen
	}

	if start < 0 {
		absStart := int(math.Abs(float64(start)))
		start = elLen - absStart

		if absStart > elLen {
			start = 0
		}
	}

	if deleteCount < 0 {
		deleteCount = start
	}

	if deleteCount > elLen - start {
		deleteCount = elLen - start
	}

	// Add item
	if itemLen > 0 {
		*elems = append(
			pElems,
			item...
		)
		sub := Slice(pElems, start, elLen)
		prefixEl := Slice(pElems, 0, start)
		addEl := Slice(*elems, elLen)

		c := make([]string, 0)
		c = append(c, prefixEl...)
		c = append(c, addEl...)
		c = append(c, sub...)
		*elems = c
	}

	// Delete item
	if deleteCount > 0 {
		endIdx := start + deleteCount
		if itemLen > 0 {
			start += itemLen
			endIdx += itemLen
			elLen += itemLen
		}
		// Shallow copy
		deleteItem := make([]string, deleteCount)
		copy(deleteItem, (*elems)[start: endIdx])
		*elems = append((*elems)[0: start], (*elems)[endIdx: elLen]...)
		return deleteItem
	}

	return []string{}
}
