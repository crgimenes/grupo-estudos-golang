package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func main() {
	// arquivos para comprimir
	files := []string{
		"teste1.txt",
		"teste2.txt",
		"teste3.txt",
	}

	// cria o arquivo .zip
	zipFile, err := os.Create("arquivo.zip")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer zipFile.Close() // quando terminar fecha o arquivo .zip

	// cria um writer que vai escrever no arquivo .zip
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close() // quando terminar fecha o writer

	for i, file := range files {

		fmt.Printf("comprimindo arquivo #%02d %v\n", i+1, file)

		// abre o arquivo que vai ser comprimido
		var fileToCompress *os.File
		fileToCompress, err = os.Open(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer fileToCompress.Close()

		// pega infos do arquivo que vai ser comprimido
		var info os.FileInfo
		info, err = fileToCompress.Stat()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("tamanho original %d bytes\n", info.Size())

		// prepara informações do arquivo que vai ser comprimido
		// para colocar no cabeçalho do arquivo .zip
		var header *zip.FileHeader
		header, err = zip.FileInfoHeader(info)
		if err != nil {
			fmt.Println(err)
			return
		}

		// ajusta metodo de compressão
		//header.Method = zip.Store
		header.Method = zip.Deflate

		// grava cabeçalho do zip
		var writer io.Writer
		writer, err = zipWriter.CreateHeader(header)
		if err != nil {
			fmt.Println(err)
			return
		}

		// grava arquivo comprimido no arquivo .zip
		_, err = io.Copy(writer, fileToCompress)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
