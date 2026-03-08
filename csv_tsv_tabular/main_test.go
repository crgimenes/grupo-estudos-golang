package main

import "testing"

func TestParseCSV(t *testing.T) {
	rows, err := ParseCSV("a,b\nc,d\n", ',')
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 2 || rows[1][1] != "d" {
		t.Fatalf("unexpected rows: %#v", rows)
	}
}

func TestParseTSV(t *testing.T) {
	rows, err := ParseCSV("a\tb\n", '\t')
	if err != nil {
		t.Fatal(err)
	}
	if rows[0][1] != "b" {
		t.Fatalf("unexpected rows: %#v", rows)
	}
}
