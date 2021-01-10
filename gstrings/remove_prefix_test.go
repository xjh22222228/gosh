package gstrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemovePrefix(t *testing.T) {
	_assert := assert.New(t)

	_assert.Equal("Hook", RemovePrefix("TestHook", "Test"))
	_assert.Equal("BaseTestCase", RemovePrefix("BaseTestCase", "Test"))
}

func BenchmarkRemovePrefix(b *testing.B) {
	for i := 0; i <b.N; i++ {
		RemovePrefix("TestHook", "Test")
	}
}
