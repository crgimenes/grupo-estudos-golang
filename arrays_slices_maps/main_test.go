package main

import "testing"

func TestCopySlice(t *testing.T) {
	s := []int{1, 2, 3}
	c := copySlice(s)
	s[0] = 99
	if c[0] != 1 {
		t.Fatalf("copy shared memory, got %d", c[0])
	}
}

func TestWordFreq(t *testing.T) {
	freq := wordFreq([]string{"a", "a", "b"})
	if freq["a"] != 2 || freq["b"] != 1 {
		t.Fatalf("unexpected freq: %#v", freq)
	}
}
