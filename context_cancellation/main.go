package main

import (
	"context"
	"fmt"
	"time"
)

func Work(ctx context.Context, d time.Duration) error {
	select {
	case <-time.After(d):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	okCtx, okCancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer okCancel()
	if err := Work(okCtx, 20*time.Millisecond); err != nil {
		fmt.Println("unexpected:", err)
	} else {
		fmt.Println("work finished without cancellation")
	}

	cancelCtx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()
	if err := Work(cancelCtx, 50*time.Millisecond); err != nil {
		fmt.Println("canceled with:", err)
	}
}
