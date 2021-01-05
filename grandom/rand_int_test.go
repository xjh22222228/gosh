package grandom

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestRandInt(t *testing.T) {
	_assert := assert.New(t)
	_assert.LessOrEqual(RandInt(5, 5), 5)
	_assert.GreaterOrEqual(RandInt(5, 100), 5)

	fmt.Println(RandInt(-10, math.MaxInt32))
}
