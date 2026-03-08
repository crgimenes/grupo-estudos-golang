package main

import (
	"strings"
	"testing"
)

func TestLogRequest(t *testing.T) {
	logger, buf := NewBufferLogger()
	LogRequest(logger, "/health", 200, 3)
	out := buf.String()
	if !strings.Contains(out, "\"path\":\"/health\"") {
		t.Fatalf("unexpected log: %s", out)
	}
}
