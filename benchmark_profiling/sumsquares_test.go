package main

import "testing"

func TestSumSquares(t *testing.T) {
	if got := SumSquares(3); got != 14 {
		t.Fatalf("SumSquares = %d", got)
	}
}

func BenchmarkSumSquares(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = SumSquares(1000)
	}
}
