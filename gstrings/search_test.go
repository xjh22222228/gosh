package gstrings

import (
	"testing"
)

func TestSearch(t *testing.T) {
	str := "The quick brown fox jumps over the lazy dog. If the dog barked, was it really lazy?"

	if v, _ := Search(str, "[^\\w\\s]"); v != 43 {
		t.Fatalf("'%v', expect: 43", v)
	}

	str2 := "hey JudE"
	if v, _ := Search(str2, "[A-Z]"); v != 4 {
		t.Fatalf("'%v', expect: 4", v)
	}

	if v, _ := Search(str2, "[.]"); v != -1 {
		t.Fatalf("'%v', expect: -1", v)
	}
}
