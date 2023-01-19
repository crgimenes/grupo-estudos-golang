package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	keyHex := flag.String("key", "", "Key")
	flag.Parse()
	key, err := hex.DecodeString(*keyHex)
	if err != nil {
		panic(err)
	}
	h := hmac.New(sha256.New, key)
	if _, err = io.Copy(h, os.Stdin); err != nil {
		panic(err)
	}
	fmt.Println("MAC:", hex.EncodeToString(h.Sum(nil)))
}
