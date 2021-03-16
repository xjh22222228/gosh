package gstr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLastIndexOf(t *testing.T) {
	_assert := assert.New(t)
	str := "canal"

	_assert.Equal(3, LastIndexOf(str, "a"))
	_assert.Equal(1, LastIndexOf(str, "a", 2))
	_assert.Equal(-1, LastIndexOf(str, "a", 0))
	_assert.Equal(-1, LastIndexOf(str, "x"))
	_assert.Equal(0, LastIndexOf(str, "c", -5))
	_assert.Equal(0, LastIndexOf(str, "c", 0))
	_assert.Equal(5, LastIndexOf(str, ""))
	_assert.Equal(2, LastIndexOf(str, "", 2))

	_assert.Equal(10, LastIndexOf("Brave new world", "w"))
	_assert.Equal(6, LastIndexOf("Brave new world", "new"))
}

