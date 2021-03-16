package gstr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrimStart(t *testing.T) {
	_assert := assert.New(t)
	str := "   Hello world!   "

	_assert.Equal("Hello world!   ", TrimStart(str))

	str = `

h`
	_assert.Equal("h", TrimStart(str))
}
