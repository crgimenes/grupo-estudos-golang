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
		leitor := bufio.NewReader(os.Stdin)
		fmt.Print("texto a ser enviado: ")
		texto, erro2 := leitor.ReadString('\n')
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
