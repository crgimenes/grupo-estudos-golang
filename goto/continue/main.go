package main

import (
	"fmt"
)

func main() {
label:
	for coluna := 10; coluna > 0; coluna-- {
		for linha := 10; linha > 0; linha-- {
			fmt.Println(linha, coluna)
			if linha == 5 {
				continue label
			}
		}
	}
}
