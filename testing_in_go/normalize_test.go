package main

import (
	"fmt"
	"testing"
)

func assertEqual(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestNormalizeSpace(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"single spaces", "go is fun", "go is fun"},
		{"extra spaces", "go   is   fun", "go is fun"},
		{"tabs and newlines", "go\t\nis\tfun", "go is fun"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assertEqual(t, NormalizeSpace(tc.in), tc.want)
		})
	}
}

func ExampleNormalizeSpace() {
	fmt.Println(NormalizeSpace("go    is\tfun"))
	// Output: go is fun
}
