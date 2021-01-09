package gslice

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestSlice(t *testing.T) {
	_assert := assert.New(t)
	animals := []string{"ant", "bison", "camel", "duck", "elephant"}

	_assert.Equal("camel,duck,elephant", strings.Join(Slice(animals, 2), ","))
	_assert.Equal("camel,duck", strings.Join(Slice(animals, 2, 4),
		","))
	_assert.Equal("bison,camel,duck,elephant", strings.Join(Slice(animals, 1, 5),
		","))

	fruits := []string{"Banana", "Orange", "Lemon", "Apple", "Mango"}
	citrus := Slice(fruits, 1, 3)
	_assert.Equal("Orange,Lemon", strings.Join(citrus, ","))
	citrus[0] = "copy"
	// Shallow copy
	_assert.Equal("Banana,Orange,Lemon,Apple,Mango", strings.Join(fruits, ","))
}

