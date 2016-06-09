# Funções

Go permite que uma função tenha mais de um retorno o que facilita bastante e torna o código mais limpo.

```
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

func main() {
	fmt.Printf("Funções!\r\n")

	// função anonima que vamos passar para printFunc
	f := func(v string) string {
		return "Olá " + v + "!\r\n"
	}

	fmt.Printf("Soma 1+1 = %v\r\n", sum(1, 1))

	r1, r2 := swap("A", "B")
	fmt.Printf("troca 1, 2 = %v, %v\r\n", r1, r2)

	printFunc(f, "Cesar")

}
```

---
[Inicio](README.md)
