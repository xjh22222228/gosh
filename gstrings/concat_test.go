package gstrings

import (
	"testing"
)

func TestConcat(t *testing.T) {
	if s := Concat("a", "b", "c"); s != "abc" {
		t.Fatalf("'%v', expect: abc", s)
	}

}
