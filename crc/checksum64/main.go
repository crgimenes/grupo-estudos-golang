package main

import (
	"fmt"
	"hash/crc64"
)

func main() {
	valor := "Isto é um teste"
	crcTable := crc64.MakeTable(crc64.ECMA)
	checksum := crc64.Checksum([]byte(valor), crcTable)
	fmt.Printf("Checksum 64 bits: 0x%x\n", checksum)
}
