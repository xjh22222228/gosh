package gstrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	_assert := assert.New(t)

	_assert.Equal("dlrow olleh", Reverse("hello world"))
}

