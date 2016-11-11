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

	potato := "Batata"
	for indice, letra := range potato {
		fmt.Printf("potato[%v] = %q\r\n", indice, letra)
	}
}
```

---
[Inicio](../README.md)

[< Funções](../funcoes/) - [switch >](../switch/)
