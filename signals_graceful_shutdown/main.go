package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func WaitForShutdown(parent context.Context, sigs <-chan os.Signal) context.Context {
	ctx, cancel := context.WithCancel(parent)
	go func() {
		select {
		case <-ctx.Done():
		case <-sigs:
			cancel()
		}
	}()
	return ctx
}

func SignalChannel() chan os.Signal {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	return ch
}
