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

func BenchmarkKeys(b *testing.B) {
    m := map[string]string{
        "g": "",
        "o": "",
        "s": "",
        "h": "",
        "!": "",
    }

    for i := 0; i < b.N; i++ {
        Keys(m)
    }
}
