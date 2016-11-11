# Olá Mundo

Se tudo estiver certo já conseguimos rodar nosso primeiro programa.

```go
package main

import "fmt"

func main() {
    fmt.Printf("Olá Mundo!\r\n")
}
```

Para rodar use o seguinte comando:

```bash
go run ola_mundo.go
```

**Lembre!** Você esta usando um compilador, compiladores alteram arquivos executáveis... antivirus não gostam que se altere arquivos executáveis e as vezes bloqueiam o compilador sem dar nenhum alerta. *Desative seu antivirus antes de programar!*

Quando usamos o comando `run`, na verdade go compila o programa para um diretório temporario. para go exibir esse diretório é só passar o parâmetro --work na linha de comando.

```bash
go run --work ola_mundo.go
```

Para compilar o programa em um executável independente use o seguinte comando:

```bash
go build ola_mundo.go
```

Agora o programa pode ser executado simplesmente sendo chamado na linha de comando e não precisa mais do ambiente do Go instalado.


## Boas praticas

Um pouco de boas praticas.

Go tem uma ferramenta de formatação de código que deixa seu código direitinho no formato padronizado... então larga a mão de ser teimoso e roda um `go fmt`.

```bash
go fmt ola_mundo.go
```


---
[Inicio](../README.md)

[< Configurando](../configurando.md) - [Variáveis >](../variaveis/)
