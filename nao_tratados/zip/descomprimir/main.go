package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func main() {
	// abre arquivo .zip
	zipFile, err := zip.OpenReader("arquivo.zip")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer zipFile.Close()

	for i, file := range zipFile.File {

		fmt.Printf("descomprimindo arquivo #%02d %v\n", i+1, file.Name)

		// abre reader para ler arquivo de dentro do zip
		reader, err := file.Open()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer reader.Close()

		var f *os.File
		// abre arquivo de destino
		f, err = os.OpenFile(
			file.Name,
			os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
			file.Mode())
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		// grava arquivo de destino
		_, err = io.Copy(f, reader)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
