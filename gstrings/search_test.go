package gstrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	_assert := assert.New(t)
	str := "The quick brown fox jumps over the lazy dog. If the dog barked, was it really lazy?"

	_assert.Equal(43, Search(str, "[^\\w\\s]"))

	str = "hey JudE"
	_assert.Equal(4, Search(str, "[A-Z]"))
	_assert.Equal(-1, Search(str, "[.]"))
}
