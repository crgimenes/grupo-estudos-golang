package main

import "testing"

func TestParseDelimitedCSV(t *testing.T) {
	rows, err := ParseDelimited("name,total\nAna,42\n", ',')
	if err != nil {
		t.Fatal(err)
	}

	want := "42"
	got := rows[1][1]
	if got != want {
		t.Fatalf("ParseDelimited()[1][1] = %q, want %q", got, want)
	}
}

func TestParseDelimitedTSV(t *testing.T) {
	rows, err := ParseDelimited("name\ttotal\nAna\t42\n", '\t')
	if err != nil {
		t.Fatal(err)
	}

	want := "Ana"
	got := rows[1][0]
	if got != want {
		t.Fatalf("ParseDelimited()[1][0] = %q, want %q", got, want)
	}
}

func TestParseDelimitedRejectsBrokenQuotes(t *testing.T) {
	_, err := ParseDelimited("name,total\nAna,\"42\n", ',')
	if err == nil {
		t.Fatal("ParseDelimited() error = nil, want parse error")
	}
}

func TestFormatRows(t *testing.T) {
	rows := [][]string{
		{"name", "city", "total"},
		{"Ana", "Santos", "42.50"},
	}

	want := "name | city | total\nAna | Santos | 42.50\n"
	got := FormatRows(rows)
	if got != want {
		t.Fatalf("FormatRows() = %q, want %q", got, want)
	}
}
