# tcp

### Trabalhando com o protocolo tcp

Go através do pacote net provê interfaces de acesso a I/O, incluindo a pilha TCP/IP, UDP, resolução por nome de domínio e UNIX Sockets. 

Neste exemplo, iremos criar um cliente TCP/IP e um servidor TCP/IP.

Servidor.

O protocolo de comunicação que nosso servidor deverá trabalhar será este:
1) Ouvir a interface tcp na porta 8081
2) Aceitar conexões
3) Dentro de um loop infinito, ouvir as mensagens a serem transmitidas pelo cliente
4) Escrever no terminal estas mensagens
5) Devolver a mensagem recebida ao cliente (eco)

```go

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	fmt.Println("Servidor aguardando conexões...")

	// ouvindo na porta 8081 via protocolo tcp/ip
	ln, erro1 := net.Listen("tcp", ":8081")
	if erro1 != nil {
		fmt.Println(erro1)
		/* Neste nosso exemplo vamos convencionar que a saída 3 está reservada para erros de conexão.
		IMPORTANTE: defers não serão executados quando utilizamos os.Exit() e a saída será imediata */
		os.Exit(3)
	}

	// aceitando conexões
	conexao, erro2 := ln.Accept()
	if erro2 != nil {
		fmt.Println(erro2)
		os.Exit(3)
	}

	defer ln.Close()

	fmt.Println("Conexão aceita...")
	// rodando loop contínuo (até que ctrl-c seja acionado)
	for {
		// Assim que receber o controle de nova linha (\n), processa a mensagem recebida
		mensagem, erro3 := bufio.NewReader(conexao).ReadString('\n')
		if erro3 != nil {
			fmt.Println(erro3)
			os.Exit(3)
		}

		// escreve no terminal a mensagem recebida
		fmt.Print("Mensagem recebida:", string(mensagem))

		// para um exemplo simples de processamento, converte a mensagem recebida para caixa alta
		novamensagem := strings.ToUpper(mensagem)

		// envia a mensagem processada de volta ao cliente
		conexao.Write([]byte(novamensagem + "\n"))
	}
}



```

Cliente.
O protocolo de comunicação que nosso cliente deverá trabalhar será este:
1) Conectar-se a interface tcp localhost na porta 8081
2) Dentro de um loop infinito, realizar leitura do terminal
3) Escrever no socket a mensagem digitada no terminal (transmitir)
4) Ouvir o retorno do servidor
5) Escrever no terminal a mensagen retornada (eco)


```go


package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {

	// conectando na porta 8081 via protocolo tcp/ip na máquina local
	conexao, erro1 := net.Dial("tcp", "127.0.0.1:8081")
	if erro1 != nil {
		fmt.Println(erro1)
		os.Exit(3)
	}

	for {
		// lendo entrada do terminal
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("texto a ser enviado: ")
		texto, erro2 := reader.ReadString('\n')
		if erro2 != nil {
			fmt.Println(erro2)
			os.Exit(3)
		}

		// escrevendo a mensagem na conexão (socket)
		fmt.Fprintf(conexao, texto+"\n")

		// ouvindo a resposta do servidor (eco)
		mensagem, err3 := bufio.NewReader(conexao).ReadString('\n')
		if err3 != nil {
			fmt.Println(err3)
			os.Exit(3)
		}

		// escrevendo a resposta do servidor no terminal
		fmt.Print("Resposta do servidor: " + mensagem)
	}
}



```

Para ter uma boa experiência deste exemplo, execute primeiramente o tcp-server e em seguida o tcp-client
preferencialmente em terminais diferentes para acompanhar o resultado.



---
[Inicio](../README.md)

[< Cliente Http](../http_templates/) - [Trabalhando com Arquivo >](../ioutil/)