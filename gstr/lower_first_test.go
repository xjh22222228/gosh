package gstr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLowerFirst(t *testing.T) {
	_assert := assert.New(t)

	_assert.Equal("gOSH", LowerFirst("GOSH"))
	_assert.Equal("", LowerFirst(""))
	_assert.Equal(" GOSH", LowerFirst(" GOSH"))
	_assert.Equal("gosh", LowerFirst("gosh"))
}
