package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	l, err := net.Listen("unix", "/tmp/echo.sock")
	if err != nil {
		panic(err)
	}

	for {
		f, err := l.Accept()
		if err != nil {
			panic(err)
		}

		go func(c io.ReadWriter) {
			for {
				buf := make([]byte, 512)
				n, err := c.Read(buf)
				if err != nil {
					return
				}

				fmt.Printf("echo: %s\n", buf[:n])
				_, err = c.Write(buf[:n])
				if err != nil {
					panic(err)
				}
			}
		}(f)
	}
}
