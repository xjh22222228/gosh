package gstrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlice(t *testing.T) {
	_assert := assert.New(t)
	str := "The quick brown fox jumps over the lazy dog."

	_assert.Equal("the lazy dog.", Slice(str, 31))
	_assert.Equal("quick brown fox", Slice(str, 4, 19))
	_assert.Equal("dog.", Slice(str, -4))
	_assert.Equal("lazy", Slice(str, -9, -5))

	str = "The morning is upon us."
	_assert.Equal("us.", Slice(str, -3))
	_assert.Equal("us", Slice(str, -3, -1))
	_assert.Equal("The morning is upon us", Slice(str, 0, -1))

	// None
	_assert.Equal("", Slice("", -3))
	_assert.Equal("1", Slice("1", -3))
	_assert.Equal("111", Slice("111", -3))
	_assert.Equal("111", Slice("1111", -3))
}
