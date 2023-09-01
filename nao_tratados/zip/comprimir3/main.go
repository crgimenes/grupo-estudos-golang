package main

import (
	"log"

	"github.com/go-br/estudos/zip/comprimir3/zip"
)

/**********************************
  Este exemplo mostra como criar um
  arquivo zip mantendo gravando
  diretamente no descritor de arquivos.

  É uma abordagem melhor que manter
  os dados todos na RAM, o unico problema
  é que o arquivo precisa ser criado antes
  de iniciar a compressão e se o sistema
  falhar pode deixar sugeira no disco
  confundir o usuário etc.

***********************************/

func main() {
	// conteudo dos arquivos
	file1 := []byte("conteudo arquivo 1")
	file2 := []byte("conteudo arquivo 2")
	file3 := []byte("conteudo arquivo 3")

	// cria instancia e adiciona
	// arquivos ao zip
	z, err := zip.New("arquivos.zip")
	if err != nil {
		log.Fatal(err)
	}
	// adiciona arquivos no zip
	err = z.Add("arquivo1.txt", file1)
	if err != nil {
		log.Fatal(err)
	}
	// para esse exemplo ignoramos
	// as mensagens de erro
	_ = z.Add("arquivo2.txt", file2)
	_ = z.Add("arquivo3.txt", file3)

	err = z.Close()
	if err != nil {
		log.Fatal(err)
	}
}
