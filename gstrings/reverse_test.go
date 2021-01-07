package gstrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	_assert := assert.New(t)

	_assert.Equal("dlrow olleh", Reverse("hello world"))
	_assert.Equal("好你界世HSOG", Reverse("GOSH世界你好"))
	_assert.Equal("", Reverse(""))
}
