package gmath

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRound(t *testing.T) {
	_assert := assert.New(t)

	_assert.Equal(4.0, Round(4.006))
	_assert.Equal(4.01, Round(4.006, 2))
}
