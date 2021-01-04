package grandom

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandInt(t *testing.T) {
	_assert := assert.New(t)
	_assert.LessOrEqual(RandInt(5, 5), 5)
	_assert.GreaterOrEqual(RandInt(5, 100), 5)
}
