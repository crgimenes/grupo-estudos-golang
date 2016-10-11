# switch


> Ao contrário das outras linguagens de programação em Go clausula switch não passa automaticamente para o próximo case, ao contrario é necessário usar *fallthrough* para que isso aconteça. É uma forma de evitar o esquecimento de colocar um *break* no fim de cada comando case.


```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("UNIX box?\r\n")
	switch os := runtime.GOOS; os {
	case "darwin":
		fallthrough
	case "freebsd":
		fallthrough
	case "openbsd":
		fallthrough
	case "plan9":
		fmt.Printf("YES!\r\n.")
	case "linux":
		fmt.Printf("almost...\r\n")
	default:
		fmt.Printf("not at all...\r\n")
	}
}

```


---
[Inicio](README.md)

[< Loop for](for.md) - [defer >](defer.md)
