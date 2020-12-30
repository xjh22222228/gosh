package gstrings

import "testing"

func TestStartsWith(t *testing.T) {
	if v := StartsWith("hello world", "hello", 0); !v {
		t.Fatalf("'%v', expect: true", v)
	}

	if v := StartsWith("hello world", "hello"); !v {
		t.Fatalf("'%v', expect: true", v)
	}

	if v := StartsWith("hello world", "hello", 1); v {
		t.Fatalf("'%v', expect: false", v)
	}
}
