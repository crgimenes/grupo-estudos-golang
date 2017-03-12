# Plugins

A possibilidade de criar plugins é um novo recurso da versão 1.8 que por enquanto funciona apenas no Linux. Um plugin basicamente é uma biblioteca externa que pode ser carregada dinamicamente, o uso é bem simples.

O mesmo recurso deve ser incluso no macOS na versão 1.9 e provavelmente em outras plataformas.

Por enquanto para todos poderem experimentar sugerimos subir um container Docker.


O comando abaixo abre um container Docker com uma maquina virtual pronta para executar a versão 1.8rc3 e monta o diretório atual.

```bash
docker run -it -v $(pwd):/go --entrypoint bash golang:1.8rc3-wheezy
``` 

Um plugin é escrito da mesma maneira que estamos acostumados, veja o exemplo abaixo.

```go
package main

import "fmt"

func Hello() error {
	fmt.Println("Hello, plugins!")
	return nil
}
```

A diferença esta apenas na compilação, um plugin deve ser compilado da seguinte maneira.

```bash
go build -buildmode=plugin
```

Isso vai criar um arquivo *plugin.so* que vai conter as funções que poderem ser carregadas por um programa principal.

Caso você queira que um plugin tenha um nome diferente de *plugin.so* basta usar o parâmetro *-o* como no exemplo:

```bash
go build -buildmode=plugin -o novonome.so
```

Para usar um plugin primeiro carregamos o arquivo.

```go
var p *plugin.Plugin
var err error

p, err = plugin.Open("./plugin/plugin.so")
if err != nil {
	fmt.Println(err)
	return
}
```

Em seguida procuramos pela função dentro do arquivo do plugin.

```go
var s plugin.Symbol

s, err = p.Lookup("Hello")
if err != nil {
	fmt.Println(err)
	return
}
```

A variável *s* do exemplo anterior é na verdade uma interface, então é necessário converter a variável para ter a mesma assinatura da função que Hello que veio do plugin. Vamos aproveitar a oportunidade para criar uma variável com o mesmo nome da função e assim poder chamar a função *sem arestas*

```go
Hello := s.(func() error)
```

A partir desse momento podemos chamar a função *Hello* normalmente como faríamos com qualquer função do nosso código.

```go
err = Hello()
if err != nil {
	fmt.Println(err)
	return
}
```

O recurso de plugins pode ser útil para criar sistemas menos monolíticos mas isso não parece mais tão relevante como teria sido no passado, hoje em dia onde tudo roda em containers autocontidos carregar uma biblioteca dinamicamente não é mais tão importante. Esse recurso também pode ser útil para distribuir binários de parte do sistema sem precisar enviar os fontes.

---
[Inicio](../README.md)

[< testing](../testing/) - [data race >](../data-race/)
