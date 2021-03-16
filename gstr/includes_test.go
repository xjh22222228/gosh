package gstr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIncludes(t *testing.T) {
	str := "To be, or not to be, that is the question."
	_assert := assert.New(t)

	_assert.True(Includes(str, "To be"))
	_assert.True(Includes(str, "question", 0))
	_assert.False(Includes(str, "nonexistent", 0))
	_assert.False(Includes(str, "To be", 1))
	_assert.False(Includes(str, "TO BE", 0))
}
