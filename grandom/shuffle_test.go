package grandom

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShuffleInt(t *testing.T) {
	_assert := assert.New(t)
	v := []int{1, 2, 3, 4, 5}
	ShuffleInt(v)

	_assert.Len(v, 5)
}

func TestShuffle(t *testing.T) {
	_assert := assert.New(t)
	v := []string{"a", "b", "c", "d", "e"}
	Shuffle(v)

	_assert.Len(v, 5)
}

func TestShuffleAny(t *testing.T) {
	_assert := assert.New(t)
	v := []interface{}{"a", false, true, 1, "e", 2, nil}
	ShuffleAny(v)
	_assert.Len(v, len(v))
}
