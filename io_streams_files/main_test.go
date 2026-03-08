package main

import (
	"errors"
	"io"
	"strings"
	"testing"
)

func TestReadFirstLine(t *testing.T) {
	got, err := ReadFirstLine(strings.NewReader("a\nb\n"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "a" {
		t.Fatalf("got %q", got)
	}
}

func TestReadFirstLineEOF(t *testing.T) {
	_, err := ReadFirstLine(strings.NewReader(""))
	if !errors.Is(err, io.EOF) {
		t.Fatalf("expected EOF, got %v", err)
	}
}
