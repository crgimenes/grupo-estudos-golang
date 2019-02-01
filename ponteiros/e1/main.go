package main

import "fmt"

type obj struct {
	value string
}

func (o *obj) update() {
	// o é um ponteiro para o objeto então value sera alterado
	o.value = "novo valor 1"
}

func (o obj) update2() {
	// o é uma copia então value não sera alterado
	o.value = "novo valor 2"
}

func main() {

	o := obj{
		value: "valor original",
	}

	fmt.Println("1", o.value)
	o.update()
	fmt.Println("2", o.value)
	o.update2()
	fmt.Println("3", o.value)

}
