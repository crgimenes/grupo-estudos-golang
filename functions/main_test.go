package main

import "testing"

func TestDivide(t *testing.T) {
	got, err := divide(8, 2)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if got != 4 {
		t.Fatalf("divide = %v, want 4", got)
	}

	_, err = divide(1, 0)
	if err == nil {
		t.Fatal("expected error for division by zero")
	}
}

func TestMakeAdder(t *testing.T) {
	add3 := makeAdder(3)
	if got := add3(4); got != 7 {
		t.Fatalf("got %d, want 7", got)
	}
}
