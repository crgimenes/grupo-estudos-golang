package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d\r\n", i)
	}
}

