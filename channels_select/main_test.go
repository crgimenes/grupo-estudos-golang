package main

import (
	"testing"
	"time"
)

func TestFirstSignalFast(t *testing.T) {
	if got := FirstSignal(200 * time.Millisecond); got != "fast" {
		t.Fatalf("got %q", got)
	}
}

func TestFirstSignalTimeout(t *testing.T) {
	if got := FirstSignal(1 * time.Millisecond); got != "timeout" {
		t.Fatalf("got %q", got)
	}
}
