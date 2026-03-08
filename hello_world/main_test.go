package main

import "testing"

func TestHelloMessage(t *testing.T) {
	want := "Hello, world!"
	got := helloMessage()
	if got != want {
		t.Fatalf("helloMessage() = %q, want %q", got, want)
	}
}
