# defer

A clausula defer define que uma função deve ser chamada no fim da execução da rotina atual.
Essa clausula é muito útil para por exemplo fechar arquivos abertos durante a execução da função.

```go
f, err := os.Open("filename.ext")
if err != nil {
    log.Fatal(err)
}
defer f.Close()

```

Mais um exemplo:

```go
package main

import (
	"fmt"
)

func main() {

	defer func() {
		fmt.Printf("Fim da função main\r\n")
	}()

	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d\r\n", i)
		fmt.Printf("dentro do for %d\r\n", i)
	}
}
```


---
[Inicio](../README.md)

[< switch](../switch/) - [interface >](../interface/)
