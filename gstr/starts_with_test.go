package gstr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStartsWith(t *testing.T) {
	_assert := assert.New(t)

	_assert.True(StartsWith("hello world", "hello", 0))
	_assert.True(StartsWith("hello world", "hello"))
	_assert.False(StartsWith("hello world", "hello", 1))
}
