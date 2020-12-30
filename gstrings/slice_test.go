package gstrings

import (
	"testing"
)

func TestSlice(t *testing.T) {
	str := "The quick brown fox jumps over the lazy dog."

	if v := Slice(str, 31); v != "the lazy dog." {
		t.Fatalf("'%v', expect: the lazy dog.", v)
	}

	if v := Slice(str, 4, 19); v != "quick brown fox" {
		t.Fatalf("'%v', expect: quick brown fox", v)
	}

	if v := Slice(str, -4, len(str)); v != "dog." {
		t.Fatalf("'%v', expect: dog.", v)
	}

	if v := Slice(str, -9, -5); v != "lazy" {
		t.Fatalf("'%v', expect: lazy", v)
	}

	str = "The morning is upon us."
	if v := Slice(str, -3); v != "us." {
		t.Fatalf("'%v', expect: us.", v)
	}
	if v := Slice(str, -3, -1); v != "us" {
		t.Fatalf("'%v', expect: us", v)
	}
	if v := Slice(str, 0, -1); v != "The morning is upon us" {
		t.Fatalf("'%v', expect: The morning is upon us", v)
	}
}
