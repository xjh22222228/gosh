package gfunctools

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersionCompare(t *testing.T) {
	_assert := assert.New(t)

	_assert.Equal(-1, VersionCompare("v0.0.1", "v0.0.2"))
	_assert.Equal(-1, VersionCompare("0.0.1", "0.0.2"))
	_assert.Equal(-1, VersionCompare("V0.0.1", "V0.0.2"))
	_assert.Equal(-1, VersionCompare("V0.0.1", "0.0.2"))
	_assert.Equal(0, VersionCompare("V0.0.2", "0.0.2"))
	_assert.Equal(1, VersionCompare("V1.0.2", "0.0.2"))
	_assert.Equal(1, VersionCompare("V1", "0.0.2"))
}

func BenchmarkVersionCompare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VersionCompare("V1", "0.0.2")
	}
}
