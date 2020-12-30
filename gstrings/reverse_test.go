package gstrings

import (
	"testing"
)

func TestReverse(t *testing.T) {
	if s := Reverse("hello world"); s != "dlrow olleh" {
		t.Fatalf("'%v', expect: dlrow olleh", s)
	}
}

