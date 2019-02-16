# shadow funções builtin

Go tem poucas funções internas (builtin) que são bastante uteis e um erro comum é declarar uma variável ou função que redefine essas funções (shadow).

Por exemplo declarar uma variável com nome de len vai sobrepor a função len(), como no exemplo:

```go
package main

import "fmt"

func main() {
    len := 10 // shadow função builtin len()
    fmt.Printf("len = %v\n", len)
}
```

Tudo vai funcionar a menos que você queira usar realmente a função len().

Links úteis com as funções internas do Go:
https://golang.org/pkg/builtin/
https://golang.org/src/builtin/builtin.go
