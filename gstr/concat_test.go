package gstr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcat(t *testing.T) {
	_assert := assert.New(t)

	_assert.Equal("abc", Concat("a", "b", "c"))
}
