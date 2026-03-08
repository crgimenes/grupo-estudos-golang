package main

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestWorkCanceled(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()

	err := Work(ctx, 100*time.Millisecond)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("err = %v", err)
	}
}

func TestWorkDone(t *testing.T) {
	ctx := context.Background()
	if err := Work(ctx, 1*time.Millisecond); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
