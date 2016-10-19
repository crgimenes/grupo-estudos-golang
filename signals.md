# Sinais

Tratar sinais é uma boa pratica, dessa forma você pode finalizar seu programa graciosamente, liberando recursos, fechando banco de dados, etc. muito melhor que simplesmente fechar o programa. E tratar sinais do sistema operacional com Go é muito simples porque basicamente o sistema vai jogar o sinal em uma canal e então precisamos apenas ficar escutando.

Primeiro criamos um canal

```go
sc := make(chan os.Signal, 1)
```

Daí informamos qual o tipo de sinal estamos interessados, no caso ^C ou seja [SIGINT](https://en.wikipedia.org/wiki/Unix_signal#SIGINT).

```go
signal.Notify(sc, os.Interrupt)
```

Então basta ficar esperando o canal retornar

```go
<-sc
```

Claro que você precisa tratar sinais em uma goroutine, dessa forma seu programa não fica parado esperando pelo sinal.

### Exemplo completo

Agora vamos ver como tudo isso trabalha junto em um [programa completo](https://github.com/crgimenes/Go-Hands-On/blob/master/signals.go).

```go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {

	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, os.Interrupt)
		// espera pelo sinal
		<-sc

		fmt.Printf("\r\nliberando recursos...\r\n")

		// Aqui você fecha o banco de dados, libera memória, etc...

		fmt.Printf("have a nice day!\r\n")
		os.Exit(0)
	}()

	fmt.Printf("Pressione ^C para terminar\r\n")

	for {
		/*
			fica colocando pontos na tela a cada segundo só
			para mostrar que o programa esta rodando
		*/
		time.Sleep(1 * time.Second)
		fmt.Printf(".")
	}

}
```

---
[Inicio](README.md)

[< http](http.md)
