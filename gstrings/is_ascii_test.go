package gstrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsAscii(t *testing.T) {
	_assert := assert.New(t)
	str := "\u0000\u0001\u0002\u0003\u0004\u0005\u0006\a\b\t\n\v\f\n\u000E" +
		"\u000F\u0010\u0011\u0012\u0013\u0014\u0015\u0016\u0017\u0018\u0019\u001A\u001B\u001C\u001D\u001E\u001F !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"

	_assert.True(IsAscii(str))

	// 
	_assert.False(IsAscii("\u0080"))

	_assert.False(IsAscii("abcd好"))
}

func BenchmarkIsAscii(b *testing.B) {
	str := "\u0000\u0001\u0002\u0003\u0004\u0005\u0006\a\b\t\n\v\f\n\u000E" +
		"\u000F\u0010\u0011\u0012\u0013\u0014\u0015\u0016\u0017\u0018\u0019\u001A\u001B\u001C\u001D\u001E\u001F !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"
	for i := 0; i < b.N; i++ {
		IsAscii(str)
	}
}
