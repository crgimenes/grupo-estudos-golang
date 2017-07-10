package main

import (
	"fmt"
)

var ( // Declarando multiplas variáveis
	nome                   = "Cesar"
	idade                  = 42
	url, path, query, page = "example.org", "/search", "&q=golang", 1
)

const pi = 3.141592

func main() {
	var x, y int = 1, 2

	var a string
	var s string

	a = "texto 1"
	b := "texto 2"

	ola := func() {
		fmt.Printf("Olá da função anônima!\r\n")
	}

	fmt.Printf("a tipo: %T\r\n", a)
	fmt.Printf("b tipo: %T\r\n", b)
	fmt.Printf("π tipo: %T\r\n", pi)
	fmt.Printf("ola tipo: %T\r\n", ola)

	fmt.Printf("valor de a = %v\r\n", a)
	fmt.Printf("valor de b = %v\r\n", b)
	fmt.Printf("valor de π = %v\r\n", pi)

	fmt.Printf("valor de x = %v\r\n", x)
	fmt.Printf("valor de y = %v\r\n", y)

	fmt.Printf("valor de s = %q\r\n", s)

	ola()
}
