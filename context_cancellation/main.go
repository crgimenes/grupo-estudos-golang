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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()
	fmt.Println(Work(ctx, 50*time.Millisecond) != nil)
}
