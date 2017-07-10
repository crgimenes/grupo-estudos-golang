package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		rt := rand.Int31n(1000)
		time.Sleep(time.Duration(rt) * time.Millisecond)
		c1 <- "retorno rotina 1"
	}()

	go func() {
		rt := rand.Int31n(1000)
		time.Sleep(time.Duration(rt) * time.Millisecond)
		c2 <- "retorno rotina 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("canal 1 retornou :", msg1)
		case msg2 := <-c2:
			fmt.Println("canal 2 retornou :", msg2)
		}
	}

	fmt.Printf("Todas as goroutines retornaram.\r\n")

}
