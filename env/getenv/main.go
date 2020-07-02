package main

import (
	"fmt"
	"os"
)

func main() {
	home := os.Getenv("HOME")
	if home == "" {
		fmt.Println("variavel não definida.")
		return
	}
	fmt.Println("home:", home)
}
