package grandom

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandom(t *testing.T) {
	_assert := assert.New(t)
	v := Random()
	_assert.Equal(v < 1, true)
	_assert.Equal(v >= 0, true)
	_assert.Equal("float64", fmt.Sprintf("%T", v))
}
