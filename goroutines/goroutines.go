package main

import "time"

func foo() {
	println("foo")
}

func main() {
	go foo()

	// experimente comentar a linha abaixo
	time.Sleep(2 * time.Second)
	println("fim")
}
