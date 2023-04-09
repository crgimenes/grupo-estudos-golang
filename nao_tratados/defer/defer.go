package main

import (
	"fmt"
)

func main() {

	defer func() {
		fmt.Printf("Fim da função main\r\n")
	}()

	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d\r\n", i)
		fmt.Printf("dentro do for %d\r\n", i)
	}
}
