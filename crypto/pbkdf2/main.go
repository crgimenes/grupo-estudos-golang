package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
)

var (
	iter = flag.Int("iter", 1, "iter")
	key  = flag.String("key", "", "password")
	salt = flag.String("salt", "", "salt")
)

func main() {
	flag.Parse()

	keyRaw := pbkdf2.Key([]byte(*key), []byte(*salt), *iter, 256/8, sha256.New)
	fmt.Printf("%x\n", keyRaw)
}
