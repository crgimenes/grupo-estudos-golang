package greetings

import "testing"

func TestHello(t *testing.T) {
	if got := Hello("go"); got != "hello, go" {
		t.Fatalf("Hello() = %q", got)
	}
}
