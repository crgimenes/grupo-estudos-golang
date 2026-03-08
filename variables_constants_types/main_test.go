package main

import "testing"

func TestDescribeLevel(t *testing.T) {
	tests := []struct {
		level LogLevel
		want  string
	}{
		{LevelDebug, "debug"},
		{LevelInfo, "info"},
		{LevelWarn, "warn"},
		{LevelError, "error"},
		{LogLevel(99), "unknown"},
	}

	for _, tc := range tests {
		got := describeLevel(tc.level)
		if got != tc.want {
			t.Fatalf("describeLevel(%v) = %q, want %q", tc.level, got, tc.want)
		}
	}
}
