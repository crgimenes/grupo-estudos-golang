package main

import (
	"fmt"
)

func main() {
	fmt.Println("1")
	goto label
	fmt.Println("2")
label:
	fmt.Println("3")
}
