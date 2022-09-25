package gstr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCapitalize(t *testing.T) {
	_assert := assert.New(t)

	_assert.Equal("好gosh", Capitalize("好GOSH"))
	_assert.Equal("G好osh", Capitalize("G好OSH"))
	_assert.Equal("Gosh", Capitalize("GOSH"))
	_assert.Equal("Gosh", Capitalize("gosh"))
	_assert.Equal("Gosh", Capitalize("gOsH"))
	_assert.Equal("", Capitalize(""))
	_assert.Equal("好", Capitalize("好"))
	_assert.Equal("好", Capitalize([]byte("好")))
}

func BenchmarkCapitalize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Capitalize("好GOSH")
	}
}
