package gstr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveSuffix(t *testing.T) {
	_assert := assert.New(t)

	_assert.Equal("Misc", RemoveSuffix("MiscTests", "Tests"))
	_assert.Equal("TmpDirMixin", RemoveSuffix("TmpDirMixin", "Tests"))
	_assert.Equal("TmpDirMixin", RemoveSuffix("TmpDirMixin", ""))
}

func BenchmarkRemoveSuffix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RemoveSuffix("MiscTests", "Tests")
	}
}
