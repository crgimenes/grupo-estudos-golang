package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("senoide")

	for i := 0.0; i < 100; i++ {
		v := math.Sin(i / 10)
		fmt.Println(v)
	}
}
