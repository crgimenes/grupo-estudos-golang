package main

import (
	"os"
	"strings"
)

func main() {

	if len(os.Args) == 1 {
		MostraTabela(GETMap())
		return
	} else if len(os.Args) == 2 {
		MontaAlfabeto(strings.ToLower(string(os.Args[1])), "")
		return
	} else {
		MontaAlfabeto(strings.ToLower(string(os.Args[1])), strings.ToUpper(string(os.Args[2])))
		return
	}
}
