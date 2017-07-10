package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("# X\tY\n")

	for x := 0.0; x < 100; x++ {
		y := math.Sin(x / 10)
		fmt.Println(x, y)
	}
}
