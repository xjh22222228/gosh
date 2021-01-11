package gmath

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMax(t *testing.T) {
	_assert := assert.New(t)
	nums := []float64{-10, 1, 2, 3, 4, 5}

	_assert.Equal(5.0, Max(nums...))
	_assert.Equal(0.0, Max())
	_assert.Equal(5.0, Max(5))
	_assert.Equal(-5.0, Max(-10, -5))
	_assert.Equal(50.0, Max(10, 20, 30, 40, 50))
}

func BenchmarkMax(b *testing.B) {
	nums := []float64{-10, 1, 2, 3, 4, 5}

	for i := 0; i < b.N; i++ {
		Max(nums...)
	}
}
