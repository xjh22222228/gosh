package gmath

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	_assert := assert.New(t)

	_assert.Equal(6, Sum(1, 2, 3))
	_assert.Equal(111, Sum([]int{10, 20, 30, 51}...))
	_assert.Equal(-5, Sum(-10, 2, 3))
	_assert.Equal(0, Sum())
}
