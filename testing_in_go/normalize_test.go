package space_test

import (
	"fmt"
	"testing"

	space "example.com/testing-in-go"
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
		{"leading and trailing", "  go is fun\n", "go is fun"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := space.NormalizeSpace(tc.in)
			assertEqual(t, got, tc.want)
		})
	}
}

func ExampleNormalizeSpace() {
	fmt.Println(space.NormalizeSpace("go    is\tfun"))
	// Output: go is fun
}
