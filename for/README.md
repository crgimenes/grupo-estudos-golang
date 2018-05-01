# for

Go só tem uma forma de loop, o for, mas em Go for é muito flexível.

```go
package main

import (
	"fmt"
)

func main() {
	valor := 0

	for i := 0; i < 10; i++ {
		valor++
		fmt.Printf("valor +1 = %v\r\n", valor)
	}

	for {
		valor--
		fmt.Printf("valor -1 = %v\r\n", valor)

		if valor == 0 {
			break
		}
	}

	//Pode-se testar uma condicao, simulando o while(true) de outras linguagens
	teste := true
	for teste {
		fmt.Println("Vamos imprimir mais uma linha nesse for...")
		valor++
		if valor == 10 {
			teste = false
		}
	}

	potato := "Batata"
	for indice, letra := range potato {
		fmt.Printf("potato[%v] = %q\r\n", indice, letra)
	}
}
```
[Playground](https://play.golang.org/p/47Yf0l5hnx)
