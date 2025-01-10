package main

import (
	"embed"
	"fmt"
)

//go:embed hello.txt
var f embed.FS

func main() {
	b, err := f.ReadFile("hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("helo.txt:", string(b))
}
