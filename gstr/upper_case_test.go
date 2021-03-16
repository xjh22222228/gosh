package gstr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpperFirstFirst(t *testing.T) {
	_assert := assert.New(t)

	_assert.Equal("Gosh", UpperFirst("gosh"))
	_assert.Equal("", UpperFirst(""))
	_assert.Equal(" GOSH", UpperFirst(" GOSH"))
	_assert.Equal("GOSH", UpperFirst("GOSH"))
}
