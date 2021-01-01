package gstrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndsWith(t *testing.T) {
	_assert := assert.New(t)
	s := "hello world"

	_assert.True(EndsWith(s, "world"))

	s = "To be, or not to be, that is the question."
	_assert.True(EndsWith(s, "to be", 19))
	_assert.False(EndsWith(s, "to be"))
}
