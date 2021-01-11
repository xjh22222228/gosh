package gmath

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMin(t *testing.T) {
	_assert := assert.New(t)
	nums := []float64{-10, 1, 2, 3, 4, 5}

	_assert.Equal(-10.0, Min(nums...))
	_assert.Equal(0.0, Min())
	_assert.Equal(5.0, Min(5))
	_assert.Equal(-10.0, Min(-10, -5))
	_assert.Equal(10.0, Min(10, 20, 30, 40, 50))
}

func BenchmarkMin(b *testing.B) {
	nums := []float64{-10, 1, 2, 3, 4, 5}

	for i := 0; i < b.N; i++ {
		Min(nums...)
	}
}
