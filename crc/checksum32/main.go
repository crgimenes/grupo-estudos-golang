package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	valor := "isto é um teste"
	checksum := crc32.ChecksumIEEE([]byte(valor))
	fmt.Printf("Checksum 32 bits: 0x%X\n", checksum)

	if checksum == 0xCF20B55 {
		fmt.Println("CRC ok!")
	} else {
		fmt.Println("CRC Falhou!")
	}

}
