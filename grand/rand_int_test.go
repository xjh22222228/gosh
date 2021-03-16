package grand

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestRandInt(t *testing.T) {
	_assert := assert.New(t)
	_assert.LessOrEqual(RandInt(5, 5), 5)
	_assert.GreaterOrEqual(RandInt(5, 100), 5)
}

func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandInt(1 << 31, int(math.Abs(1 << 31)))
	}
}
