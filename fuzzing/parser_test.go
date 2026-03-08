package main

import "testing"

func TestParseCSVLine(t *testing.T) {
	fields, err := ParseCSVLine("a,b,c")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(fields) != 3 {
		t.Fatalf("len = %d, want 3", len(fields))
	}
}

func FuzzParseCSVLine(f *testing.F) {
	seeds := []string{"a,b,c", "", "\"a,b\",c", "x;y;z"}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(t *testing.T, input string) {
		_, _ = ParseCSVLine(input)
	})
}
