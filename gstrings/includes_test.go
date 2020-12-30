package gstrings

import "testing"

func TestIncludes(t *testing.T) {
	str := "To be, or not to be, that is the question."

	if v := Includes(str, "To be"); !v {
		t.Fatalf("'%v', expect: true", v)
	}

	if v := Includes(str, "question", 0); !v {
		t.Fatalf("'%v', expect: true", v)
	}

	if v := Includes(str, "nonexistent", 0); v {
		t.Fatalf("'%v', expect: false", v)
	}

	if v := Includes(str, "To be", 1); v {
		t.Fatalf("'%v', expect: false", v)
	}

	if v := Includes(str, "TO BE", 0); v {
		t.Fatalf("'%v', expect: false", v)
	}
}
