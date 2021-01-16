package grandom

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestColor(t *testing.T) {
	_assert := assert.New(t)

	// RGB
	_assert.True(strings.HasPrefix(Color(C_RGB), "rgb("))
	_assert.Equal(3, len(strings.Split(Color(C_RGB), ",")))

	// RGBA
	_assert.True(strings.HasPrefix(Color(C_RGBA), "rgba("))
	_assert.Equal(4, len(strings.Split(Color(C_RGBA), ",")))

	// HEX
	_assert.Equal(7, len(Color(C_HEX)))
	_assert.True(strings.HasPrefix(Color(C_HEX), "#"))
}

func BenchmarkColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Color(C_RGB)
		Color(C_RGBA)
		Color(C_HEX)
		Color(C_COLOR)
	}
}
