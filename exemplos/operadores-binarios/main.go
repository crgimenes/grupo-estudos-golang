package main

import "fmt"

type bin uint8

func (b bin) String() string {
	return fmt.Sprintf("%04b-%04b", (b >> 4), b&0x0F)
}

func main() {
	var v bin = 60 // 0011 1100
	var b bin = 13 // 0000 1101
	fmt.Printf("valor de v = %3d em binario %v\n", v, v)
	c := v & b // 0000 1100
	fmt.Printf("valor de c = %3d em binario %v | %v AND         %v = %v\n", c, c, v, b, c)
	c = v | b // 0011 1101
	fmt.Printf("valor de c = %3d em binario %v | %v OR          %v = %v\n", c, c, v, b, c)
	c = v &^ b
	fmt.Printf("valor de c = %3d em binario %v | %v NOT AND     %v = %v\n", c, c, v, b, c)
	c = v << 2
	fmt.Printf("valor de c = %3d em binario %v | %v LEFT SHIFT  %v = %v\n", c, c, v, b, c)
	c = v >> 2
	fmt.Printf("valor de c = %3d em binario %v | %v RIGHT SHIFT %v = %v\n", c, c, v, b, c)
	c = v ^ b
	fmt.Printf("valor de c = %3d em binario %v | %v XOR         %v = %v\n", c, c, v, b, c)

}
