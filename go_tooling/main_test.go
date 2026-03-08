package main

import "testing"

func TestAdd(t *testing.T) {
	if got := add(2, 3); got != 5 {
		t.Fatalf("add() = %d, want 5", got)
	}
}
