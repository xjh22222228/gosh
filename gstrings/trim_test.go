package gstrings

import (
	"testing"
)

func TestTrim(t *testing.T) {
	if v := Trim("   Hello world!   "); v != "Hello world!" {
		t.Fatalf("'%v', expect: Hello world!", v)
	}

	str := `

abc
`
	if v := Trim(str); v != "abc" {
		t.Fatalf("'%v', expect: abc", v)
	}
}
