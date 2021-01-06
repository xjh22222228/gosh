package grandom

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestChoice(t *testing.T) {
	_assert := assert.New(t)
	str := "GOSH你好，世界"
	v := Choice(str)

	_assert.True(strings.ContainsAny(str, v))
	_assert.Equal(1, utf8.RuneCountInString(v))

	_assert.Equal("G", Choice("G"))
	_assert.Equal("谢", Choice("谢"))
}