# Olá Mundo

Se tudo estiver certo já conseguimos rodar nosso primeiro programa.

```
package main

import "fmt"

func main() {
    fmt.Printf("Olá Mundo!\r\n")
}
```

Para rodar use o seguinte comando:

```
go run ola_mundo.go
```

Quando usamos o comando run, na verdade go compila o programa para um diretório de trabalho. para go exibir esse diretório é só passar o parâmetro --work na linha de comando.

```
go run --work ola_mundo.go
```

Para compilar o programa em um executável independente use o seguinte comando:

```
go build ola_mundo.go
```

Agora o programa pode ser executado simplesmente sendo chamado na linha de comando e não precisa mais do ambiente do Go instalado.


## Boas praticas

Um pouco de boas praticas.
Go tem uma ferramenta de formatação de código que deixa seu código direitinho no formato padronizado... então larga a mão de ser teimoso e roda um go fmt.

```
go fmt ola_mundo.go
```


---
[Inicio](README.md)
