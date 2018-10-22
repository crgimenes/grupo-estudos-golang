package main

import (
	"fmt"
)

func main() {
	fmt.Println("1")
	/*
		esse goto não vai compilar porque goto
		não pode pular sobre declaração de variaveis
	*/
	goto label
	var i = 1
	fmt.Println("2", i)
label:
	fmt.Println("3")
}
