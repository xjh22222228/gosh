package gstrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrimEnd(t *testing.T) {
	_assert := assert.New(t)
	str := "   Hello world!   "

	_assert.Equal("   Hello world!", TrimEnd(str))

	_assert.Equal(15, len(TrimEnd(str)))

	str = `gosh  

`
	_assert.Equal(4, len(TrimEnd(str)))
	_assert.Equal("gosh", TrimEnd(str))
}
