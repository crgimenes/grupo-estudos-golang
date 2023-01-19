package main

import (
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"golang.org/x/crypto/curve25519"
	"golang.org/x/crypto/ed25519"
	"io"
	"log"
	"os"
)

var (
	key    = flag.String("key", "", "Private key.")
	keygen = flag.Bool("gen", false, "Generate ed25519 asymmetric keypair.")
	public = flag.String("pub", "", "Remote's side Public key.")
)

func main() {
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage of", os.Args[0]+":")
		flag.PrintDefaults()
		os.Exit(2)
	}

	var privatekey ed25519.PrivateKey

	if *keygen {
		_, _ = io.ReadFull(rand.Reader, privatekey[:])

		var privateKey *[32]byte
		var publicKey *[32]byte

		privateKey, publicKey, _ = GenerateKey()

		fmt.Printf("Private= %x\n", *privateKey)
		fmt.Printf("Public= %x\n", *publicKey)
		os.Exit(0)
	}

	privatekey, err := hex.DecodeString(*key)
	if err != nil {
		log.Fatal(err)
	}
	if len(privatekey) != 32 {
		log.Fatal("curve25519: bad private key length.")
		os.Exit(1)
	}
	publickey, err := hex.DecodeString(*public)
	if err != nil {
		log.Fatal(err)
	}

	var privateKey [32]byte
	copy(privateKey[:], []byte(privatekey))
	var publicKey [32]byte
	copy(publicKey[:], []byte(publickey))

	var secret []byte
	secret = GenerateSharedSecret(privateKey, publicKey)

	fmt.Printf("Shared= %x\n", secret)
	os.Exit(0)
}

func GenerateKey() (privateKey *[32]byte, publicKey *[32]byte, err error) {
	var pub, priv [32]byte

	_, err = io.ReadFull(rand.Reader, priv[:])
	if err != nil {
		return nil, nil, err
	}

	priv[0] &= 248
	priv[31] &= 127
	priv[31] |= 64

	curve25519.ScalarBaseMult(&pub, &priv)

	return &priv, &pub, nil
}

func GenerateSharedSecret(priv, pub [32]byte) []byte {
	var secret [32]byte

	curve25519.ScalarMult(&secret, &priv, &pub)

	return secret[:]
}
