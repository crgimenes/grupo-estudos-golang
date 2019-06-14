package main

import "fmt"

func main() {

	// valor original
	var valor = "teste de ponteiro"

	// declara um ponteiro
	var p *string

	// pega o endereço de valor e coloca em p
	p = &valor

	// resolve p e altera o endereço de valor
	*p = "teste"

	fmt.Println(valor)
}
