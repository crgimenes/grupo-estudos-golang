package main

import (
	"fmt"
	"io"
	"os"
)

func closer(a io.Closer) {
	err := a.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	file, err := os.Open("nome.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer closer(file)

	file2, err := os.Open("nome.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer closer(file2)
	// ....
}
