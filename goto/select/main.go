package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	timeout := time.After(2 * time.Second)

	messages := make(chan string)
	go func() { messages <- "ping" }()

loop:
	select {
	case <-timeout:
		fmt.Println("timeout")
		os.Exit(1)

	case m := <-messages:
		fmt.Println(m)
		goto loop
	}

}
