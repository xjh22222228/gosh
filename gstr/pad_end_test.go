package gstr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPadEnd(t *testing.T) {
	_assert := assert.New(t)

	_assert.Equal("abc       ", PadEnd("abc", 10))
	_assert.Equal("abcfoofoof", PadEnd("abc", 10, "foo"))
	_assert.Equal("abc123", PadEnd("abc", 6, "123456"))
	_assert.Equal("abc", PadEnd("abc", 1))
}
