package grand

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestSample(t *testing.T) {
	_assert := assert.New(t)
	elems := []string{"g", "o", "s", "h", "你", "好", "G", "o", "L", "a", "n",
		"g"}
	v := Sample(elems, 5)

	_assert.True(strings.ContainsAny(strings.Join(elems, ""), strings.Join(v,
		"")))
	_assert.Equal(5, len(v))

	elems = []string{"g", "o", "s", "h"}
	v = Sample(elems, 4)
	_assert.Equal("gosh", strings.Join(elems, ""))
}
