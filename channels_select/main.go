package main

import (
	"fmt"
	"time"
)

func FirstSignal(timeout time.Duration) string {
	fast := make(chan string, 1)
	slow := make(chan string, 1)

	go func() {
		time.Sleep(10 * time.Millisecond)
		fast <- "fast"
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		slow <- "slow"
	}()

	select {
	case msg := <-fast:
		return msg
	case msg := <-slow:
		return msg
	case <-time.After(timeout):
		return "timeout"
	}
}

func main() {
	fmt.Println(FirstSignal(100 * time.Millisecond))
}
