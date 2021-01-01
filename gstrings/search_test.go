package gstrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	_assert := assert.New(t)
	str := "The quick brown fox jumps over the lazy dog. If the dog barked, was it really lazy?"

	v, _ := Search(str, "[^\\w\\s]")
	_assert.Equal(43, v)

	str = "hey JudE"
	v, _ = Search(str, "[A-Z]")
	_assert.Equal(4, v)

	v, _ = Search(str, "[.]")
	_assert.Equal(-1, v)
}
