package main

import (
	"log"

	"github.com/go-br/estudos/zip/comprimir2/zip"
)

/**********************************
  Este exemplo mostra como criar um
  arquivo zip mantendo usando um
  package que mantém o buffer com os
  dados todos na RAM.

  É uma abordagem interessante em alguns
  casos mas a RAM não é infinita, então
  essa abordagem só serve se temos certeza
  que os dados sempre serão peucos (sempre
	estamos errados quanto a isso)

***********************************/

func main() {
	// conteudo dos arquivos
	file1 := []byte("conteudo arquivo 1")
	file2 := []byte("conteudo arquivo 2")
	file3 := []byte("conteudo arquivo 3")

	// cria instancia e adiciona
	// arquivos ao zip
	z := zip.New()
	err := z.Add("arquivo1.txt", file1)
	if err != nil {
		log.Fatal(err)
	}
	// para esse exemplo ignoramos
	// as mensagens de erro
	_ = z.Add("arquivo2.txt", file2)
	_ = z.Add("arquivo3.txt", file3)

	err = z.Save("arquivos.zip")
	if err != nil {
		log.Fatal(err)
	}
}
