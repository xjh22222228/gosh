package gstrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCapitalize(t *testing.T) {
	_assert := assert.New(t)

	_assert.Equal("Gosh", Capitalize("GOSH"))
	_assert.Equal("Gosh", Capitalize("gosh"))
	_assert.Equal("Gosh", Capitalize("gOsH"))
	_assert.Equal("", Capitalize(""))
}
