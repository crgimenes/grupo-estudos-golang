package main

import "testing"

func TestClassify(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{-1, "negative"},
		{0, "zero"},
		{2, "even"},
		{3, "odd"},
	}
	for _, tc := range tests {
		if got := classify(tc.n); got != tc.want {
			t.Fatalf("classify(%d) = %q, want %q", tc.n, got, tc.want)
		}
	}
}

func TestSum(t *testing.T) {
	if got := sum([]int{1, 2, 3}); got != 6 {
		t.Fatalf("sum = %d, want 6", got)
	}
}
