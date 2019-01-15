package main

import (
	"fmt"
	"os"
)

func main() {
	if _, err := os.Stat("teste.txt"); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Arquivo não existe")
			// então podemos criar o arquivo...
		} else {
			fmt.Println(err)
			return
			// alguma outro erro aconteceu.
		}
	}
	// continua o programa...
}
