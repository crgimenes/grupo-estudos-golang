# Variáveis

Go é fortemente tipada, isso significa que os tipos das variáveis são bem definidos e não podem mudar durante o programa.

Tipos de variáveis em Go:

bool
string

Varias formas de inteiro:
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte na verdade é um alias para uint8

rune é um alias int32 e é usada para armazenar Unicode

float32 float64

complex64 complex128


```
package main

import (
	"fmt"
)

var ( // Declarando múltiplas variáveis
	nome       string = "Cesar"
	valor      int    = 42
	a1, a2, a3        = 10, 20, 30
)

const pi = 3.141592

func main() {
	var x,y int = 1,2

	var a string
	a = "texto 1"
	b := "texto 2"

	ola := func() {
		fmt.Printf("Olá da função anônima!\r\n")
	}

	fmt.Printf("a tipo: %T\r\n", a)
	fmt.Printf("b tipo: %T\r\n", b)
	fmt.Printf("π tipo: %T\r\n", pi)
	fmt.Printf("ola tipo: %T\r\n", ola)

	fmt.Printf("valor de a = %v\r\n", a)
	fmt.Printf("valor de b = %v\r\n", b)
	fmt.Printf("valor de π = %v\r\n", pi)

	fmt.Printf("valor de x = %v\r\n", x)
	fmt.Printf("valor de y = %v\r\n", y)

	fmt.Printf("valor de s = %q\r\n", s)

	ola()
}
```


---
[Inicio](README.md) | [Olá Mundo <-](ola_mundo.md) | [Funções ->](funcoes.md)
