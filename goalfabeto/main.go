package main

import (
        "github.com/paulopraxedes/estudos/go/goalfabeto"
        "os"
        "strings"
)

func main() {

        if len(os.Args) == 1 {
                goalfabeto.MostraTabela(goalfabeto.GETMap())
                return
        } else if len(os.Args) == 2 {
                goalfabeto.MontaAlfabeto(strings.ToLower(string(os.Args[1])), "")
                return
        } else {
                goalfabeto.MontaAlfabeto(strings.ToLower(string(os.Args[1])), strings.ToUpper(string(os.Args[2])))
                return
        }
}

