package main

import (
	"crypto/sha256"
	"fmt"
	"hash/crc32"
	"hash/crc64"
)

var crc64Table = crc64.MakeTable(crc64.ISO)

func CRC32(data []byte) uint32 {
	return crc32.ChecksumIEEE(data)
}

func CRC64(data []byte) uint64 {
	return crc64.Checksum(data, crc64Table)
}

func SHA256(data []byte) string {
	sum := sha256.Sum256(data)
	return fmt.Sprintf("%x", sum)
}

func main() {
	payload := []byte("invoice:42:paid")

	fmt.Printf("crc32  %08x\n", CRC32(payload))
	fmt.Printf("crc64  %016x\n", CRC64(payload))
	fmt.Printf("sha256 %s\n", SHA256(payload))
}
