package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	count int
	mutex sync.RWMutex
)

func printCount(label string) {
	for {
		mutex.RLock()
		fmt.Println(label, count)
		mutex.RUnlock()
		<-time.After(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}

func incCount() {
	for {
		mutex.Lock()
		count++
		mutex.Unlock()
		<-time.After(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano() )
	go printCount("goroutine 1 leu:")
	go printCount("goroutine 2 leu:")
	go printCount("goroutine 3 leu:")
	go printCount("goroutine 4 leu:")
	go incCount()

	<-time.After(5 * time.Second)
}
