# for

Go só tem uma forma de loop, o for, mas em Go for é muito flexível.

```
package main

import (
	"fmt"
)

func main() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum++
		fmt.Printf("soma +1 = %v\r\n", sum)
	}

	for {
		sum--
		fmt.Printf("soma -1 = %v\r\n", sum)

		if sum == 0 {
			break
		}
	}

	potato := "Batata"
	for a, b := range potato {
		fmt.Printf("potato[%v] = %q\r\n", a, b)
	}
}
```

---
[Inicio](README.md) | [Variáveis <-](funcoes.md) | [switch ->](switch.md)
