package main

import (
	"errors"
	"strings"
	"testing"
)

func TestDescribe(t *testing.T) {
	if got := describe("x"); got != "x" {
		t.Fatalf("got %q", got)
	}
	err := errors.New("boom")
	if got := describe(err); got != "boom" {
		t.Fatalf("got %q", got)
	}
}

func TestReadAll(t *testing.T) {
	got, err := readAll(strings.NewReader("abc"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "abc" {
		t.Fatalf("got %q", got)
	}
}
