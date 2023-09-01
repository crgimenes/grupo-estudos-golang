package main

import "fmt"

type SuaStruct struct {
	Valor int
}

func (s SuaStruct) String() string {
	return fmt.Sprintf("esta é a sua struct e o valor é %v", s.Valor)
}

func main() {
	sua := SuaStruct{
		Valor: 10,
	}
	fmt.Println(sua)
}
