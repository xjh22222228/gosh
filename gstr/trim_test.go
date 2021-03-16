package gstr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrim(t *testing.T) {
	_assert := assert.New(t)

	_assert.Equal("Hello world!", Trim("   Hello world!   "))

	str := `

abc
`
	_assert.Equal("abc", Trim(str))
}
