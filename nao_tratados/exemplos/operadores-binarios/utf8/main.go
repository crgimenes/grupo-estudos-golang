package main

import "fmt"

func main() {
	x := "isto é um teste!"
	for i := 0; i < len(x); i++ {
		c := x[i]
		b := c >> 7
		if b == 1 {
			fmt.Printf("%08b <- UTF-8\n", c)
			continue
		}
		fmt.Printf("%08b = \"%v\"\n", c, string(c))
	}
}
