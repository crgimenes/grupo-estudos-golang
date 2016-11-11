package main

import (
	"log"
	"os"
)

func main() {

	log.Print("Tentando abrir arquivo\r\n")
	f, err := os.Open("filename.ext")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	log.Print("O arquivo abriu sem erro.\r\n")
}
