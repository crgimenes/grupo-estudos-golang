# Olá, mundo!

Este tópico apresenta a estrutura mínima de um programa Go e os comandos essenciais para executar e compilar.

## Objetivo

Entender os elementos básicos de um programa Go:

- `package main`;
- função `main`;
- execução com `go run`;
- compilação com `go build`.

## Estrutura do exemplo

- `main.go`: programa mínimo com saída no terminal.
- `main_test.go`: teste unitário da função que gera a mensagem.

## Código

```go
package main

import "fmt"

func helloMessage() string {
	return "Hello, world!"
}

func main() {
	fmt.Println(helloMessage())
}
```

## Executar

No diretório `hello_world`:

```bash
go run .
```

Saída esperada:

```text
Hello, world!
```

## Compilar

```bash
go build .
```

O comando gera um executável no diretório atual.

## Testar

```bash
go test ./...
```

## Pontos importantes

- Em Go, o programa executável começa em `package main`.
- A entrada do programa sempre é a função `main`.
- `go run` compila e executa temporariamente.
- `go build` gera binário para execução posterior.
