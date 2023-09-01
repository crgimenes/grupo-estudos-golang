package main

import (
	_ "embed"
	"fmt"
)

//go:embed "hello.txt"
var s string

func main() {
	fmt.Println("helo.txt:", s)
}
