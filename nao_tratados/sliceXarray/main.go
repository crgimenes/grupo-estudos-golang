package main

import (
	"fmt"
)

func main() {
	// este é um array
	a := [6]string{"isso", "é", "um", "coleção", "de", "palavras"}

	// aqui criamos um slice apontando para o array a
	s := a[:]

	fmt.Printf("valor de a[3]: %q\n", a[3])
	fmt.Printf("valor de s[3]: %q\n", s[3])

	// como s é um ponteiro para a, ao mudar
	// s também mudamos a
	fmt.Println(`-- Como "s" aponte para "a", mudar "s" também muda "a"`)
	s[3] = "array"
	fmt.Printf("valor de a[3]: %q\n", a[3])
	fmt.Printf("valor de s[3]: %q\n", s[3])

	// vamos agora fazer um append no slice s
	fmt.Println(`-- Fazer append em "s" cria um novo array`)
	s = append(s, "!")
	// isso criou um novo array e copiou os dados antigos
	// agora vamos mudar o valor de s[3] para demonstrar isso
	s[3] = "slice"

	fmt.Printf("valor de a[3]: %q\n", a[3])
	fmt.Printf("valor de s[3]: %q\n", s[3])

}
