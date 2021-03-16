package gstr

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unicode/utf8"
)

func TestShuffle(t *testing.T) {
	_assert := assert.New(t)

	_assert.Equal(16, utf8.RuneCountInString(Shuffle("世界你好 hello world")))
}
