package gstr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCenter(t *testing.T) {
	_assert := assert.New(t)

	_assert.Equal("[gosh]", Center("[gosh]", 0))
	_assert.Equal("[gosh]", Center("[gosh]", 4))
	_assert.Equal("------------[gosh]------------", Center("[gosh]", 30, '-'))
	_assert.Equal("------------[gosh]-----------", Center("[gosh]", 29, '-'))
	_assert.Equal("好好好好好好好好好好好好[gosh]好好好好好好好好好好好", Center("[gosh]", 29, '好'))
}
