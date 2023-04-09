package main

import "fmt"

func main() {
	len := 10 // shadow função builtin len()
	fmt.Printf("len = %v\n", len)
}
