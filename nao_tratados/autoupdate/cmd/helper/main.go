package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var mutex sync.Mutex

func write(msg string) {
	mutex.Lock()
	os.Stdout.Write([]byte(msg))
	mutex.Unlock()
}

func main() {
	go func() {
		for {
			<-time.After(1 * time.Second)
			write("ping\n")
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("ready\n")
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		s := scanner.Text()
		log.Printf("launcher: %q", s)
		switch s {
		case "quit":
			write("bye\n")
			os.Exit(0)

		case "ping":
			write("pong\n")
		}
	}
}
