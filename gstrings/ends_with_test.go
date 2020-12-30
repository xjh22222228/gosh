package gstrings

import (
	"testing"
)

func TestEndsWith(t *testing.T) {
	s := "hello world"

	if v := EndsWith(s, "world"); !v {
		t.Fatalf("'%v', expect: true", v)
	}

	str := "To be, or not to be, that is the question."
	if v := EndsWith(str, "to be", 19); !v {
		t.Fatalf("'%v', expect: true", v)
	}

	if v := EndsWith(str, "to be", len(str)); v {
		t.Fatalf("'%v', expect: false", v)
	}
}
