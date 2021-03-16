package gstr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndexOf(t *testing.T) {
	_assert := assert.New(t)
	str := "hello world"

	_assert.Equal(0, IndexOf(str, ""))
	_assert.Equal(0, IndexOf(str, "", 0))
	_assert.Equal(3, IndexOf(str, "", 3))
	_assert.Equal(8, IndexOf(str, "", 8))
	_assert.Equal(11, IndexOf(str, "", 11))
	_assert.Equal(11, IndexOf(str, "", 13))
	_assert.Equal(11, IndexOf(str, "", 22))

	str = "Blue Whale"
	_assert.Equal(0, IndexOf(str, "Blue"))
	_assert.Equal(-1, IndexOf(str, "Blute"))
	_assert.Equal(5, IndexOf(str, "Whale", 0))
	_assert.Equal(5, IndexOf(str, "Whale", 5))
	_assert.Equal(0, IndexOf(str, "", -1))
	_assert.Equal(9, IndexOf(str, "", 9))
	_assert.Equal(10, IndexOf(str, "", 10))
	_assert.Equal(10, IndexOf(str, "", 11))
	_assert.Equal(-1, IndexOf(str, "blue"))
}
