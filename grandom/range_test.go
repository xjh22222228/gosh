package grandom

import (
	"github.com/stretchr/testify/assert"
	"github.com/xjh22222228/gosh/gmath"
	"testing"
)

func TestRange(t *testing.T) {
	_assert := assert.New(t)
	_assert.Equal(4, len(Range(0, 10, 3)))
	_assert.Equal(0, len(Range(5, 5)))
	_assert.Equal(0, len(Range(5, 5, 2)))
	_assert.Equal(5050, gmath.Sum(Range(0, 100)...))
	_assert.Equal(18, gmath.Sum(Range(0, 10, 3)...))
}

func BenchmarkRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Range(0, 100000, 3)
	}
}
