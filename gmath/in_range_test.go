package gmath

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInRange(t *testing.T) {
	_assert := assert.New(t)

	_assert.True(InRange(3, 2, 4))
	_assert.True(InRange(-3, -2, -6))
	_assert.True(InRange(4, 8))
	_assert.True(InRange(1.2, 2))
	_assert.False(InRange(4, 2))
	_assert.False(InRange(2, 2))
	_assert.False(InRange(5.2, 4))
}
