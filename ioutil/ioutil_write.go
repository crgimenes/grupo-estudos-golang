package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

//Função para tratamento de erro
func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func main() {

	//Texto a ser gravado
	file1 := []byte("Go \n Hands-On \n")

	// Parm 1: nome do artquivo
	// Parm 2: conteúdo do arquivo
	// Parm 3: permissão de gravação ver os.FileMode
	err1 := ioutil.WriteFile("file1", file1, 0644)
	check(err1)

	// Apenas criando um novo arquivo vazio
	f, err2 := os.Create("file2")
	check(err2)
	// Feche imediatamente assim que finalizar a função principal.
	defer f.Close()

	// slice de bytes a ser gravada
	file2 := []byte{103, 111, 45, 104, 97, 110, 100, 115, 45, 111, 110}
	// neste caso não definimos o nome do arquivo a ser gravado
	// e a biblioteca utilizou o nome do arquivo do programa principal "ioutil_write"
	n1, err3 := f.Write(file2)
	check(err3)
	// n1 recebe o total de bytes gravados
	fmt.Printf("%s \n", string(file2))
	fmt.Printf("gravados %d bytes \n", n1)

	// Cria o arquivo
	file3, err4 := os.Create("file3")
	check(err4)
	// Emita uma sincronização para liberar gravações para um armazenamento estável.
	// O pacote Sync fornece primitivas básicas de sincronização, tal como bloqueios de exclusão mútua (lock).
	file3.Sync()
	// bufferiza
	w := bufio.NewWriter(file3)
	n2, err5 := w.WriteString("Go-Hands-On bufferizado")
	check(err5)
	fmt.Printf("gravados %d bytes \n", n2)
	// Flush garante que todos os dados foram encaminhados serão tratados pelo io.Writer.
	w.Flush()
	file3.Close()

	// Abre o arquivo com a opção de append (acrescentar)
	file, err6 := os.OpenFile("file3", os.O_APPEND|os.O_WRONLY, 0644)
	// os.O_WRONLY - Essas solicitações abrem o arquivo somente leitura, somente gravação ou leitura / gravação, respectivamente.
	check(err6)
	n3, err7 := file.WriteString("Go-Hands-On acrescentando")
	check(err7)
	fmt.Printf("gravados %d bytes \n", n3)
	file.Close()

}
