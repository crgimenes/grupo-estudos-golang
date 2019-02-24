package main

import "fmt"

type comptometer interface {
	Sum(a, b int) int
}

type foo struct{}

func (_ foo) Sum(a, b int) int {
	return a + b
}

type bar struct{}

func (_ bar) Sum(a, b int) int {
	return a + b
}

//
func printType(e comptometer) {
	/*
		Este switch esta na ordem errada
		e vai imprimir sempre "comptometer
		interface" e nunca via entrar nos
		outros cases do switch. Mova o case
		comptometer para o fim do switch
		para resolver o bug.
	*/
	switch e.(type) {
	case comptometer:
		fmt.Println("comptometer interface")
	case *foo:
		fmt.Println("ponteiro para foo")
	case *bar:
		fmt.Println("ponteiro para bar")
	}
}

func main() {
	var f = &foo{}
	printType(f)
	var b = &bar{}
	printType(b)
}
