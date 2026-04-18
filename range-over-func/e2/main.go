package main

import (
	"fmt"
	"iter"
)

func Lines(data []string) iter.Seq[string] {
	return func(yield func(string) bool) {

		for _, s := range data {
			done := func() bool {
				fmt.Println("--- Recurso Aberto ---")
				defer fmt.Println("--- Recurso Fechado via Defer ---")
				if !yield(s) {
					// Se yield retorna false, o chamador deu 'break'.
					fmt.Println("Interrupção detectada no yield!")
					return true
				}
				return false
			}()
			if done {
				break
			}
		}
	}
}

func main() {
	items := []string{"A", "B", "C", "D"}

	fmt.Println("Execução 1: Loop completo")
	for s := range Lines(items) {
		fmt.Printf("Recebido: %s\n", s)
	}

	fmt.Println("\nExecução 2: Break se receber 'B'")
	for s := range Lines(items) {
		fmt.Printf("Recebido: %s\n", s)
		if s == "B" {
			fmt.Println("Sai do loop")
			break
		}
	}
}
