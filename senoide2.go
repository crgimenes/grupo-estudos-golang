package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("# X\tY\n")

	for i := 1.0; i < 40; i += 0.01 {
		v := math.Sin(i) * (i / 2.0 * math.Pi)
		fmt.Printf("%v\t%v\n", i, v)
	}
}
