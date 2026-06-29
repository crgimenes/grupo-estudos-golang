package main

import (
	"reflect"
	"testing"
)

func TestParseKeyValues(t *testing.T) {
	in := "a=1\ninvalid\n# comment\n\nb=2\n"
	got, err := ParseKeyValues(in)
	if err != nil {
		t.Fatalf("ParseKeyValues() error = %v", err)
	}

	want := map[string]string{
		"a": "1",
		"b": "2",
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ParseKeyValues() = %#v, want %#v", got, want)
	}
}
