package main

import (
	"fmt"
	"runtime"
	"strconv"
)

func main() {
	fmt.Print("UNIX box?\r\n")
	switch runtime.GOOS {
	case "darwin":
		fallthrough
	case "freebsd":
		fallthrough
	case "openbsd":
		fallthrough
	case "plan9":
		fmt.Printf("YES!\r\n")
	case "linux":
		fmt.Printf("almost...\r\n")
	default:
		fmt.Printf("not at all...\r\n")
	}

	fmt.Println("Checando números de 1 a 10\r")
	fmt.Print("Digite um número: ")
	var inserido string
	fmt.Scanln(&inserido)

	numero, _ := strconv.Atoi(inserido)

	switch numero {
	case 1, 3, 5, 7:
		fmt.Printf("%v é primo!\n\r", numero)
		fallthrough
	case 9:
		fmt.Printf("%v é impar!\n\r", numero)
		fmt.Printf("O resto da divisão de %v por 2 é %v\n\r", numero, numero%2)
	case 2, 4, 6, 8, 10:
		fmt.Printf("%v é par!\n\r", numero)
	default:
		fmt.Printf("%v não esta entre 1 e 10!\n\r", numero)
	}

	//switch com condicional
	fmt.Println("Checando dia de receber o salário...\r")
	switch {
	case numero < 5:
		fmt.Println("Ainda não recebi meu salário.")
	case numero == 5:
		fmt.Println("Opa! Hoje é o dia de receber o money.")
	case numero > 5 && numero <= 15:
		fmt.Println("O dinheiro está acabando...!")
	case numero == 30 || numero == 31:
		fmt.Println("Xiii o money acabou.")
	default:
		fmt.Println("Parabéns! Você é rico.")
	}

	// fmt.Println(numero)
	// ./switch.go:44: undefined: numero
}
