package main

import (
	"fmt"
	"os"
)

func main() {
	home, ok := os.LookupEnv("HOME")
	if !ok {
		fmt.Println("variavel não definida.")
		return
	}
	fmt.Println("home:", home)
}
