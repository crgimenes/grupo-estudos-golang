package main

import "fmt"

func main() {
	fmt.Print("Nome: ")
	var nome string
	_, err := fmt.Scanln(&nome)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Olá,", nome)
}
