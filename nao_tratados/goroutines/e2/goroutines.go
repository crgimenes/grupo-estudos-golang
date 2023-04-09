package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printValue(value int, c chan int) {
	defer fmt.Printf("%v terminou\r\n", value)
	rt := rand.Int31n(1000)
	time.Sleep(time.Duration(rt) * time.Millisecond)
	//fmt.Printf("printValue: valor = %v\r\n", value)
	c <- value
}

func main() {
	rand.Seed(time.Now().Unix())

	c := make(chan int)

	for i := 0; i < 10; i++ {
		go printValue(i, c)
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("main: %vº colocado, valor = %v\r\n", i+1, <-c)
	}

}
