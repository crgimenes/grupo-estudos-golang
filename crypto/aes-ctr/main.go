package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	keyHex := flag.String("key", "", "Key")
	flag.Parse()
	var key []byte
	var err error
	if *keyHex == "" {
		key = make([]byte, 256/8)
		_, err = io.ReadFull(rand.Reader, key)
		if err != nil {
			panic(err)
		}
		fmt.Fprintln(os.Stderr, "Key:", hex.EncodeToString(key))
	} else {
		key, err = hex.DecodeString(*keyHex)
		if err != nil {
			panic(err)
		}
		if len(key) != 256/8 {
			panic(errors.New("provided key has wrong length"))
		}
	}
	ciph, _ := aes.NewCipher(key)
	iv := make([]byte, aes.BlockSize)
	stream := cipher.NewCTR(ciph, iv)
	buf := make([]byte, 128*1<<10)
	var n int
	for {
		n, err = os.Stdin.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		stream.XORKeyStream(buf[:n], buf[:n])
		if _, err := os.Stdout.Write(buf[:n]); err != nil {
			panic(err)
		}
		if err == io.EOF {
			break
		}
	}
}
