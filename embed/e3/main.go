package main

import (
	"embed"
	"fmt"
)

//go:embed files/* assets/*
//go:embed html/index.html
var f embed.FS

func main() {
	b, err := f.ReadFile("files/hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("helo.txt:", string(b))

	d, err := f.ReadDir("files")
	if err != nil {
		fmt.Println(err)
		return
	}
	for k, v := range d {
		fmt.Println(k, v)
	}
}
