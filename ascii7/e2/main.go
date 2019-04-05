package main

import "fmt"

type bin uint8

func (b bin) String() string {
	return fmt.Sprintf("%04b-%04b", (b >> 4), b&0x0F)
}

func stringToBitArray(s string) []bin {
	var b []bin
	for i := 0; i < len(s); i++ {
		b = append(b, bin(s[i]))
	}
	return b
}

// GSM A
func toASCII7(in []bin) (out []bin) {
	var (
		count uint
		q     uint
	)
	for _, b := range in {
		q |= (uint(b) & 0x7F) << count
		count += 7
		if count >= 8 {
			out = append(out, bin(q))
			count -= 8
			q >>= 8
		}
	}
	if count > 0 {
		out = append(out, bin(q))
	}
	return
}

func main() {
	s := "teste de string"
	v := stringToBitArray(s)
	if len(v)%7 != 0 {
		fmt.Println("não divisivel por 7")
	}
	for k, b := range v {
		fmt.Printf("%02d = %v = %02X = \"%c\"\n", k, b, uint(b), b)
	}

	fmt.Println("-=-=-=-=-=-=-=-=-=-")
	x := toASCII7(v)
	for k, b := range x {
		fmt.Printf("%02d = %v = %02X\n", k, b, uint(b))
	}
	fmt.Println("-=-=-=-=-=-=-=-=-=-")
	fmt.Println("string......:", s)
	fmt.Print("ASCII 8 bits: ")
	for _, b := range v {
		fmt.Printf("%02X", uint(b))
	}
	fmt.Print("\nASCII 7 bits: ")
	for _, b := range x {
		fmt.Printf("%02X", uint(b))
	}
	println()
}
