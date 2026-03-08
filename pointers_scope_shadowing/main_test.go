package main

import "testing"

func TestIncrement(t *testing.T) {
	v := 1
	increment(&v)
	if v != 2 {
		t.Fatalf("v = %d, want 2", v)
	}
}

func TestParseShadowing(t *testing.T) {
	got, err := parse(true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 10 {
		t.Fatalf("got %d, want 10", got)
	}
}
