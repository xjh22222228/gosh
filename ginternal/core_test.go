package ginternal

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestAddZero(t *testing.T) {
    _assert := assert.New(t)
    _assert.Equal("01", AddZero("1"))
    _assert.Equal("10", AddZero("10"))
}
