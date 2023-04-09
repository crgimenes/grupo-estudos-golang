package main

import (
	"log"

	"github.com/tarm/serial"
)

func main() {
	c := &serial.Config{Name: "/dev/cu.usbmodemFA131", Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	//n, err := s.Write([]byte("test"))
	//if err != nil {
	//	log.Fatal(err)
	//}

	buf := make([]byte, 128)
	n, err := s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q", buf[:n])
	err = s.Close()
	if err != nil {
		log.Fatal(err)
	}
}
