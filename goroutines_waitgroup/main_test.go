package main

import "testing"

func TestConcurrentSquare(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   []int
	}{
		{
			name:   "keeps input order",
			values: []int{2, 3, 4},
			want:   []int{4, 9, 16},
		},
		{
			name:   "empty input",
			values: nil,
			want:   []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ConcurrentSquare(test.values)
			if len(got) != len(test.want) {
				t.Fatalf("got len=%d want=%d", len(got), len(test.want))
			}

			for i := range test.want {
				if got[i] != test.want[i] {
					t.Fatalf("got[%d]=%d want=%d", i, got[i], test.want[i])
				}
			}
		})
	}
}
