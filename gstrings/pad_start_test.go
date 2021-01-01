package gstrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPadStart(t *testing.T) {
	str := "abc"
	_assert := assert.New(t)

	_assert.Equal("       abc", PadStart(str, 10))
	_assert.Equal("foofoofabc", PadStart(str, 10, "foo"))
	_assert.Equal("123abc", PadStart(str, 6, "123465"))
	_assert.Equal("00000abc", PadStart(str, 8, "0"))
	_assert.Equal("abc", PadStart(str, 1))
}
