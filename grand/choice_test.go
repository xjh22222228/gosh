package grand

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

func TestChoiceInt(t *testing.T) {
	_assert := assert.New(t)

	_assert.Equal(1, ChoiceInt([]int{1}))
	_assert.Equal(1, ChoiceInt([]int{1, 1, 1, 1, 1, 1}))
}

func TestChoiceStr(t *testing.T) {
	_assert := assert.New(t)

	_assert.Equal("G", ChoiceStr([]string{"G"}))
	_assert.Equal("G", ChoiceStr([]string{"G", "G", "G", "G", "G", "G"}))
	_assert.Equal("你", ChoiceStr([]string{"你", "你", "你", "你", "你"}))
}