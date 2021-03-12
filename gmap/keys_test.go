package gmap

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestKeys(t *testing.T) {
    _assert := assert.New(t)

    m := map[string]string{
        "g": "",
        "o": "",
        "s": "",
        "h": "",
        "!": "",
    }
    v := Keys(m)

    _assert.Equal(5, len(v))
}
