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

	//varrendo um map por chave e valor com for
	meal := map[string]string{"1": "Arroz", "2": "Feijão", "3": "Batata"}

	//exemplo 1, varrendo map por chave
	for key := range meal {
		fmt.Println("Chave Map: ", key)
	}

	//exemplo 2, varrendo map por valor
	for _, value := range meal {
		fmt.Println("Valor Map: ", value)
	}

	//exemplo 3, varrendo map por chave e valor
	for key, value := range meal {
		fmt.Printf("Chave: %s - Valor: %s\n", key, value)
	}
}
```
[Playground](https://play.golang.org/p/5qGMl5sY_io)
