package main

import (
	"fmt"
)

func main() {
	items := []string{"A", "B", "C", "D"}

	fmt.Println("Execução 1: Loop completo")
	for _, s := range items {
		func() {
			fmt.Println("--- Recurso Aberto ---")
			defer fmt.Println("--- Recurso Fechado via Defer ---")

			fmt.Printf("Recebido: %s\n", s)
		}()
	}

	fmt.Println("\nExecução 2: Break se receber 'B'")
	for _, s := range items {
		done := func() bool {
			fmt.Println("--- Recurso Aberto ---")
			defer fmt.Println("--- Recurso Fechado via Defer ---")

			fmt.Printf("Recebido: %s\n", s)

			if s == "B" {
				fmt.Println("Sai do loop...")
				return true
			}
			return false
		}()
		if done {
			break
		}
	}
}
