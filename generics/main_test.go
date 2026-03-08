package main

import "testing"

func TestMin(t *testing.T) {
	if got := Min(4, 2); got != 2 {
		t.Fatalf("Min = %d", got)
	}
}

func TestMapSlice(t *testing.T) {
	got := MapSlice([]int{1, 2, 3}, func(v int) int { return v * 2 })
	want := []int{2, 4, 6}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got[%d]=%d want=%d", i, got[i], want[i])
		}
	}
}
