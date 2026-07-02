package main

import (
	"slices"
	"testing"
)

func TestGzipRoundtrip(t *testing.T) {
	in := []byte("go study group\n")

	compressed, err := GzipBytes(in)
	if err != nil {
		t.Fatal(err)
	}

	out, err := UngzipBytes(compressed)
	if err != nil {
		t.Fatal(err)
	}

	if string(out) != string(in) {
		t.Fatalf("got %q, want %q", out, in)
	}
}

func TestZipSingleFile(t *testing.T) {
	archive, err := ZipSingleFile("note.txt", []byte("go study group\n"))
	if err != nil {
		t.Fatal(err)
	}

	names, err := ZipFileNames(archive)
	if err != nil {
		t.Fatal(err)
	}

	want := []string{"note.txt"}
	if !slices.Equal(names, want) {
		t.Fatalf("got %v, want %v", names, want)
	}
}
