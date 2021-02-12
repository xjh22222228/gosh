package gmath

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestClamp(t *testing.T) {
    _assert := assert.New(t)

    _assert.Equal(-5, Clamp(-10, -5, 5))
    _assert.Equal(5, Clamp(10, -5, 5))
}

func TestClampFloat64(t *testing.T) {
    _assert := assert.New(t)

    _assert.Equal(-5.0, ClampFloat(-10, -5, 5))
    _assert.Equal(5.0, ClampFloat(10, -5, 5))
}
