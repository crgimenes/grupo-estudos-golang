package main

import "testing"

func TestSafeDivideRecover(t *testing.T) {
	_, err := safeDivide(10, 0)
	if err == nil {
		t.Fatal("expected error from recovered panic")
	}
}

func TestDeferOrder(t *testing.T) {
	got := deferOrder()
	want := "start,first,second,third"
	if got != want {
		t.Fatalf("deferOrder() = %q, want %q", got, want)
	}
}
