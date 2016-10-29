package main

import (
	"fmt"
)

// Retorno simples
func sum(x int, y int) int {
	return x + y
}

// Retorno duplo
func swap(x string, y string) (string, string) {
	return y, x
}

// Retorno assinado
func div(x, y int) (result int) {
	result = x / y
	return
}

// função que recebe uma função como parâmetro
func printFunc(f func(string) string, valor string) {
	aux := f(valor)
	fmt.Printf(aux)
}

func main() {
	fmt.Printf("Funções!\r\n")

	fmt.Printf("Soma 1+1 = %v\r\n", sum(1, 1))

	r1, r2 := swap("A", "B")
	fmt.Printf("troca A, B = %v, %v\r\n", r1, r2)

	resultadoDivisao := div(10, 5)
	fmt.Printf("divisao 10 / 5 = %v\r\n", resultadoDivisao)

	// função anonima que vamos passar para printFunc
	f := func(v string) string {
		return "Olá " + v + "!\r\n"
	}

	printFunc(f, "Cesar")

}
