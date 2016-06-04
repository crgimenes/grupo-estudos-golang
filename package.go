package main

import (
	"fmt"
	r "github.com/crgimenes/rotateString"
)

func main() {
	s := r.Rotate("Isto é um teste de Rotate")
	fmt.Printf("rotate: %s\r\n", s)
}
