package hello

import "testing"

func TestHello(t *testing.T) {
	h := New()
	h.Print()
	h.Free()
}
