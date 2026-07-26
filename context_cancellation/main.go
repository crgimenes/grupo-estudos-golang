package main

import (
	"context"
	"fmt"
	"time"
)

func Work(ctx context.Context, d time.Duration) error {
	timer := time.NewTimer(d)
	defer timer.Stop()

	select {
	case <-timer.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	okCtx, okCancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer okCancel()

	err := Work(okCtx, 20*time.Millisecond)
	if err != nil {
		fmt.Println("unexpected:", err)
		return
	}
	fmt.Println("work finished without cancellation")

	cancelCtx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()

	err = Work(cancelCtx, 50*time.Millisecond)
	if err != nil {
		fmt.Println("canceled with:", err)
	}
}
