package gstrings

import "testing"

func TestCenter(t *testing.T) {
	if v := Center("[gosh]", 0); v != "[gosh]" {
		t.Fatalf("'%v', expect: [gosh]", v)
	}

	if v := Center("[gosh]", 4); v != "[gosh]" {
		t.Fatalf("'%v', expect: [gosh]", v)
	}

	if v := Center("[gosh]", 30, '-'); v != "------------[gosh]------------" {
		t.Fatalf("'%v', expect: ------------[gosh]------------", v)
	}

	if v := Center("[gosh]", 29, '-'); v != "------------[gosh]-----------" {
		t.Fatalf("'%v', expect: ------------[gosh]------------", v)
	}

	if v := Center("[gosh]", 29, '好'); v != "好好好好好好好好好好好好[gosh]好好好好好好好好好好好" {
		t.Fatalf("'%v', expect: 好好好好好好好好好好好好[gosh]好好好好好好好好好好好", v)
	}
}
