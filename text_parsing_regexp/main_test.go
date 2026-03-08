package main

import "testing"

func TestParseKeyValues(t *testing.T) {
	in := "a=1\ninvalid\nb=2\n"
	got := ParseKeyValues(in)
	if got["a"] != "1" || got["b"] != "2" {
		t.Fatalf("unexpected map: %#v", got)
	}
}
