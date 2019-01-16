package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func(ctx context.Context, cancel context.CancelFunc) {
		scInterrupt := make(chan os.Signal, 1)
		scKill := make(chan os.Signal, 1)
		signal.Notify(scInterrupt, os.Interrupt)
		signal.Notify(scKill, os.Kill)

		select {
		case <-scInterrupt:
			fmt.Println("morreu com Interrupt")
		case <-scKill:
			fmt.Println("morreu com Kill")
		}

		fmt.Println("shutting down...")

		cancel()

		//os.Exit(0)
	}(ctx, cancel)

	c := make(chan string)
	go func(ctx context.Context, cStr chan string) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println(ctx.Err())
				return
			default:
			}
			<-time.After(time.Second)
		}
		cStr <- "alguma string"
	}(ctx, c)

	select {
	case str := <-c:
		fmt.Println(str)
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
