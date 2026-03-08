package main

import (
	"context"
	"os"
	"syscall"
	"testing"
	"time"
)

func TestWaitForShutdown(t *testing.T) {
	sigs := make(chan os.Signal, 1)
	ctx := WaitForShutdown(context.Background(), sigs)
	sigs <- syscall.SIGTERM
	select {
	case <-ctx.Done():
	case <-time.After(100 * time.Millisecond):
		t.Fatal("context was not canceled")
	}
}
