package main

import "testing"

func TestConcurrentSquare(t *testing.T) {
	got := ConcurrentSquare([]int{2, 3, 4})
	want := []int{4, 9, 16}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got[%d]=%d want=%d", i, got[i], want[i])
		}
	}
}
