# Funções

* Retorno duplo - Facilita a escrita e torna o código mais limpo
* Retorno assinado - Facilita a leitura do código
* Passando valor por referencia

```go
package main

import (
	"fmt"
)

func sum(x int, y int) int {
	return x + y
}

func swap(x string, y string) (string, string) {
	return y, x
}

// função que recebe uma função como parâmetro
func printFunc(f func(string) string, valor string) {
	aux := f(valor)
	fmt.Printf(aux)
}

func printValorByRef(valor *string) {
	fmt.Printf("Valor por referencia = %v\r\n", *valor)
}

func main() {
	fmt.Printf("Funções!\r\n")

	fmt.Printf("Soma 1+1 = %v\r\n", sum(1, 1))

	r1, r2 := swap("A", "B")
	fmt.Printf("troca A, B = %v, %v\r\n", r1, r2)

	// função anonima que vamos passar para printFunc
	f := func(v string) string {
		return "Olá " + v + "!\r\n"
	}

	printFunc(f, "Cesar")

	valor := "Esse valor não vai ser copiado, só estamos passando o ponteiro"
	printValorByRef(&valor)

}
```

---
[Inicio](README.md)

[< Struct](struct.md) - [Loop for >](for.md)
