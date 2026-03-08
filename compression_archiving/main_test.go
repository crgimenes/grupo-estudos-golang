package main

import "testing"

func TestGzipRoundtrip(t *testing.T) {
	in := []byte("hello")
	z, err := GzipBytes(in)
	if err != nil {
		t.Fatal(err)
	}
	out, err := UngzipBytes(z)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "hello" {
		t.Fatalf("got %q", out)
	}
}

func TestZipSingleFile(t *testing.T) {
	z, err := ZipSingleFile("a.txt", []byte("x"))
	if err != nil {
		t.Fatal(err)
	}
	if len(z) == 0 {
		t.Fatal("empty zip")
	}
}
