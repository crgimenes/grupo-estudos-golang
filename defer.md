# defer

A clausula defer define que uma função deve ser chamada no fim da execução da rotina atual.
Essa clausula é muito útil para por exemplo fechar arquivos abertos durante a execução da função.


```
package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d\r\n", i)
	}
}
```

---
[Inicio](README.md)
