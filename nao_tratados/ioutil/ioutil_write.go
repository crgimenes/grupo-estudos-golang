package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//Texto a ser gravado
	arquivo1 := []byte("Go \n Hands-On \n")

	/* EXEMPLO 1 */
	// Parm 1: nome do artquivo
	// Parm 2: conteúdo do arquivo
	// Parm 3: permissão de gravação ver os.FileMode
	erro1 := os.WriteFile("arquivo1", arquivo1, 0644)
	if erro1 != nil {
		fmt.Println(erro1)
		/* Neste nosso exemplo vamos convencionar que a saída 3
		está reservada para erros de leitura e escrita em arquivos.
		IMPORTANTE: defers não serão executados quando utilizamos os.Exit() e a saída será imediata */
		os.Exit(3)
	}

	/* EXEMPLO 2 */
	// Apenas criando um novo arquivo vazio
	arquivo2, erro2 := os.Create("arquivo2")
	if erro2 != nil {
		fmt.Println(erro2)
		os.Exit(3)
	}

	// Feche imediatamente assim que finalizar a função principal.
	defer arquivo2.Close()

	// slice de bytes a ser gravada
	dados := []byte{103, 111, 45, 104, 97, 110, 100, 115, 45, 111, 110}
	// neste caso não definimos o nome do arquivo a ser gravado
	// e a biblioteca utilizou o nome do arquivo do programa principal "ioutil_write"
	total1, erro3 := arquivo2.Write(dados)
	if erro3 != nil {
		fmt.Println(erro3)
		os.Exit(3)
	}

	// total1 recebe o total de bytes gravados
	fmt.Printf("%s \n", string(dados))
	fmt.Printf("gravados %d bytes \n", total1)

	/* EXEMPLO 3 */
	// Cria o arquivo
	arquivo3, erro4 := os.Create("arquivo3")
	if erro4 != nil {
		fmt.Println(erro4)
		os.Exit(3)
	}

	// Emita uma sincronização para liberar gravações para um armazenamento estável.
	// O pacote Sync fornece primitivas básicas de sincronização, tal como bloqueios de exclusão mútua (lock).
	arquivo3.Sync()
	// bufferiza
	w := bufio.NewWriter(arquivo3)
	total2, erro5 := w.WriteString("Go-Hands-On bufferizado \n")
	if erro5 != nil {
		fmt.Println(erro5)
		os.Exit(3)
	}

	fmt.Printf("gravados %d bytes \n", total2)
	// Flush garante que todos os dados foram encaminhados serão tratados pelo io.Writer.
	w.Flush()
	arquivo3.Close()

	/* EXEMPLO 4 */
	// Abre o arquivo com a opção de append (acrescentar)
	arquivo, erro6 := os.OpenFile("arquivo3", os.O_APPEND|os.O_WRONLY, 0644)
	/* os.O_WRONLY - Essas solicitações abrem o arquivo somente leitura, somente gravação
	ou leitura / gravação, respectivamente. */
	if erro6 != nil {
		fmt.Println(erro6)
		os.Exit(3)
	}

	total3, erro7 := arquivo.WriteString("Go-Hands-On acrescentando")
	if erro7 != nil {
		fmt.Println(erro7)
		os.Exit(3)
	}

	fmt.Printf("gravados %d bytes \n", total3)
	arquivo.Close()

}
