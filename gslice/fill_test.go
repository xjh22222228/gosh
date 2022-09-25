package gslice

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestFill(t *testing.T) {
	_assert := assert.New(t)
	elems := []string{"1", "2", "3", "4"}
	v := Fill(elems, "0", 2, 4)

	_assert.Equal("1,2,0,0", strings.Join(v, ","))
	_assert.Equal("1,2,0,0", strings.Join(elems, ","))

	elems = []string{"1", "2", "3", "4"}
	v = Fill(elems, "5", 1)
	_assert.Equal("1,5,5,5", strings.Join(v, ","))
	_assert.Equal("1,5,5,5", strings.Join(elems, ","))

	Fill(elems, "6", 0)
	_assert.Equal("6666", strings.Join(elems, ""))

	elemInt := []int{1, 2, 3, 4}
	Fill(elemInt, 5, 1)
	_assert.Equal(5, elemInt[1])
}
