package main

import "fmt"

//go:generate go run ./cmd/gen

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Printf("%s sum=%d platform=%s\n", generatedMessage, add(2, 3), platformName())
}
