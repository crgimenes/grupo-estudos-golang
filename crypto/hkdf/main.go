package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"golang.org/x/crypto/hkdf"
	"hash"
	"io"
	"log"
)

var (
	info = flag.String("info", "", "additional Info")
	key  = flag.String("key", "", "master key")
	salt = flag.String("salt", "", "salt")
)

func main() {
	flag.Parse()

	hash, err := Hkdf([]byte(*key), []byte(*salt), []byte(*info))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x\n", hash[:256/8])
}

func Hkdf(master, salt, info []byte) ([128]byte, error) {
	var myHash func() hash.Hash
	myHash = sha256.New

	hkdf := hkdf.New(myHash, master, salt, info)

	key := make([]byte, 256/8)
	_, err := io.ReadFull(hkdf, key)

	var result [128]byte
	copy(result[:], key)

	return result, err
}
