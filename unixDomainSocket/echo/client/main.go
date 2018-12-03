package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	f, err := net.Dial("unix", "/tmp/echo.sock")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	go func(r io.Reader) {
		buf := make([]byte, 1024)
		for {
			n, errf := r.Read(buf)
			if errf != nil {
				panic(err)
			}
			fmt.Printf("recebido: %s\n", buf[:n])
		}
	}(f)

	for {
		data := []byte("olá mundo")
		fmt.Printf("enviando: %s\n", data)
		_, err = f.Write([]byte("olá mundo"))
		if err != nil {
			panic(err)
		}

		time.Sleep(time.Duration(400) * time.Millisecond)
	}
}
