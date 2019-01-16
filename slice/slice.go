package main

import "fmt"

func main() {
	//declarando básica de um slice
	var s = []int{1, 2, 3}
	fmt.Println(s)

	//declarando slice utilizando o new
	var s2 = new([6]int)[0:3]
	fmt.Printf("Tamanho s2: %v - Capacidade s2: %v \n", len(s2), cap(s2))

	//declarando slice utilizando a função make
	var s3 = make([]int, 10)     //tamanho e capacidade são 10
	var s4 = make([]int, 10, 20) //tamanho é 10 e capacidade é 20
	fmt.Printf("Tamanho s3: %v - Capacidade s3: %v \n", len(s3), cap(s3))
	fmt.Printf("Tamanho s4: %v - Capacidade s3: %v \n", len(s4), cap(s4))

	//adicionando elementos ao slice utilizando append
	s5 := make([]int, 10)
	s5 = append(s5, 1, 2, 3, 4, 5)
	fmt.Println("s5: ", s5)
	fmt.Printf("Tamanho s5: %v - Capacidade s5: %v \n", len(s5), cap(s5))

	//acessando elementos do slice
	s6 := make([]int, 3)

	s6[0] = 1
	s6[1] = 2
	s6[2] = 3

	fmt.Println(s6[0])
	fmt.Println(s6[1])
	fmt.Println(s6[:2])
	fmt.Println(s6[1:2])
	fmt.Println(s6[0:2])
	fmt.Println(s6[2:])

	//adicionando slice em outro slice
	s7 := []string{"Arroz", "Feijão", "Batata"}
	s8 := []string{"Melão", "Maçã"}

	s7 = append(s7, s8...)
	fmt.Println(s7)

}
