package gfunctools

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCurrency(t *testing.T) {
	_assert := assert.New(t)

	_assert.Equal("$123.00", Currency(123))
	_assert.Equal("$1,234.12", Currency(1234.12))
	_assert.Equal("$12,345.56", Currency(12345.56))
	_assert.Equal("$123,456.57", Currency(123456.567))
	_assert.Equal("￥1,234,567.00", Currency(1234567, "￥"))
	_assert.Equal("$12,345,678.00", Currency(12345678))
	_assert.Equal("$123,456,789.00", Currency(123456789))
}

func BenchmarkCurrency(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Currency(123456789.123456789)
	}
}
