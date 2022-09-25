package gslice

import (
	"github.com/stretchr/testify/assert"
	"math"
	"strings"
	"testing"
)

func TestSplice(t *testing.T) {
	_assert := assert.New(t)
	// Test Delete
	elems := []string{"angel", "clown", "drum", "mandarin", "sturgeon"}
	v := Splice[string](&elems, 3, 1)
	_assert.Equal("angel,clown,drum,sturgeon", strings.Join(elems, ","))
	_assert.Equal("mandarin", strings.Join(v, ""))

	elems = []string{"angel", "clown", "drum", "mandarin", "sturgeon"}
	v = Splice[string](&elems, 2, 1)
	_assert.Equal("drum", strings.Join(v, ""))

	elems = []string{"angel", "clown", "mandarin", "sturgeon"}
	v = Splice[string](&elems, -2, 1)
	_assert.Equal("mandarin", strings.Join(v, ""))

	elems = []string{"parrot", "anemone", "blue", "trumpet", "sturgeon"}
	v = Splice[string](&elems, len(elems)-3, 2)
	_assert.Equal("parrot,anemone,sturgeon", strings.Join(elems, ","))
	_assert.Equal("blue,trumpet", strings.Join(v, ","))

	// Delete All
	elems = []string{"angel", "clown", "mandarin", "sturgeon"}
	v = Splice[string](&elems, 2, math.MaxInt64)
	_assert.Equal("mandarin,sturgeon", strings.Join(v, ","))

	// Test Add
	elems = []string{"angel", "clown", "mandarin", "sturgeon"}
	v = Splice(&elems, 2, 0, "drum")
	_assert.Equal("angel,clown,drum,mandarin,sturgeon", strings.Join(elems, ","))
	_assert.Equal("", strings.Join(v, ","))

	elems = []string{"angel", "clown", "mandarin", "sturgeon"}
	v = Splice[string](&elems, 2, 0, "drum", "guitar")
	_assert.Equal("angel,clown,drum,guitar,mandarin,sturgeon",
		strings.Join(elems, ","))
	_assert.Equal("", strings.Join(v, ","))

	elems = []string{"angel", "clown", "trumpet", "sturgeon"}
	v = Splice[string](&elems, 0, 2, "parrot", "anemone", "blue")
	_assert.Equal("parrot,anemone,blue,trumpet,sturgeon",
		strings.Join(elems, ","))
	_assert.Equal("angel,clown", strings.Join(v, ","))
}

func BenchmarkSplice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		elems := []string{"angel", "clown", "drum", "mandarin", "sturgeon"}
		Splice[string](&elems, 2, 3, "a", "b", "c")
	}
}
