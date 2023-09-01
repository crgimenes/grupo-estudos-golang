package main

import "fmt"

type obj struct {
	value string
}

// muda o valor do objeto passado por referecia
func update(o *obj) {
	o.value = "novo valor obj"
}

func byRef1(v *string) {
	aux := "novo valor"

	// nesse caso não muda o valor original
	// porque mudamos o endereço de v para o
	// endereço de aux
	v = &aux
	fmt.Printf("valor na função byRref %q\n", *v)
}

func byRef2(v *string) {
	// nesse caso o valor original sera alterado
	// porque resolvemos o endereço do ponteiro e
	// alteramos o valor apontado por esse endereço
	*v = "novo valor"
	fmt.Printf("valor na função byRref %q\n", *v)
}

func byVal(v string) {
	// nunca altera o valor original porque v é uma copia
	v = "novo valor"
	fmt.Printf("valor na função byVal %q\n", v)
}

func main() {
	value := "essa é a string original"
	fmt.Printf("valor inicial %q\n", value)
	byVal(value)
	fmt.Printf("valor depois de chamar a função byVal %q\n", value)
	byRef1(&value)
	fmt.Printf("valor depois de chamar a função byRef1 %q\n", value)
	byRef2(&value)
	fmt.Printf("valor depois de chamar a função byRef2 %q\n", value)

	o := obj{
		value: "valor obj",
	}
	update(&o)
	fmt.Printf("valor depois de chamar a função update %q\n", o.value)
}
