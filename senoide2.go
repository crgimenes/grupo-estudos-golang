package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("# X\tY\n")

	for x := 1.0; x < 40; x += 0.01 {
		y := math.Sin(x) * (x / 2.0 * math.Pi)
		fmt.Printf("%v\t%v\n", x, y)
	}
}
