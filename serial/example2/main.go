package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"crg.eti.br/go/config"
	"github.com/tarm/serial"
)

type config struct {
	Port string
	Baud int `cfgDefault:"115200"`
}

func main() {

	cfg := &config{}

	err := goconfig.Parse(cfg)
	if err != nil {
		log.Fatal(err)
	}

	if cfg.Port == "" {
		fmt.Println("Porta não definida.")
	}

	c := &serial.Config{Name: cfg.Port, Baud: cfg.Baud}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, os.Interrupt)
		// espera pelo sinal
		<-sc

		fmt.Printf("\r\nliberando recursos...\r\n")
		err = s.Close()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("have a nice day!\r\n")
		os.Exit(0)
	}()

	for {
		buf := make([]byte, 128)
		n, err := s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(string(buf[:n]))
	}
}
