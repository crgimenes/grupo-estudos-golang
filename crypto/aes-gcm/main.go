package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	aad = flag.String("a", "", "additional associated data")
	dec = flag.Bool("d", false, "decrypt instead of encrypt")
	key = flag.String("k", "", "symmetric key")
)

func main() {
	flag.Parse()

	var keyHex string
	keyHex = *key
	var key []byte
	var err error
	if keyHex == "" {
		key = make([]byte, 256/8)
		_, err = io.ReadFull(rand.Reader, key)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(os.Stderr, "Key=", hex.EncodeToString(key))
	} else {
		key, err = hex.DecodeString(keyHex)
		if err != nil {
			log.Fatal(err)
		}
		if len(key) != 256/8 {
			log.Fatal(err)
		}
	}

	buf := bytes.NewBuffer(nil)
	var data io.Reader
	data = os.Stdin
	io.Copy(buf, data)
	msg := buf.Bytes()

	var c cipher.Block
	c, err = aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}
	aead, _ := cipher.NewGCMWithTagSize(c, 16)

	if *dec == false {
		nonce := make([]byte, aead.NonceSize(), aead.NonceSize()+len(msg)+aead.Overhead())

		if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
			log.Fatal(err)
		}

		out := aead.Seal(nonce, nonce, msg, []byte(*aad))
		fmt.Printf("%s", out)

		os.Exit(0)
	}

	if *dec == true {
		nonce, msg := msg[:aead.NonceSize()], msg[aead.NonceSize():]

		out, err := aead.Open(nil, nonce, msg, []byte(*aad))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", out)

		os.Exit(0)
	}
}
