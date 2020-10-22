# Arrays

Exemplos de array em Golang


```go
package main

import "fmt"

func main() {
	//declarando um array
	var a [3]int

	a[0] = 1
	a[1] = 2
	a[2] = 3

	fmt.Println(a)

	//declarando um array usando forma de declaração curta
	a2 := [3]int{4, 5, 6}
	fmt.Println(a2)

	//declarando array com valores preenchidos
	var a3 = [3]int{7, 8, 9}
	fmt.Println(a3)

	//declarando array array especificando as posições dos elementos
	var a4 = [3]int{2: 12, 1: 11, 0: 10}
	fmt.Println(a4)

	/*declarando array com elipse, onde quem determina o tamanho do array
	é a quantidade de elementos informados
	*/
	a5 := [...]int{13, 14, 15}
	fmt.Println(a5)

}
```
[Playground](https://play.golang.org/p/YmmfIIFO_By)
