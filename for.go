package main

import (
	"fmt"
)

func main() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum++
		fmt.Printf("soma +1 = %v", sum)
	}
}
