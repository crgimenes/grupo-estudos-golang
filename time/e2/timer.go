package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {

	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, os.Interrupt)
		// espera pelo sinal
		<-sc

		fmt.Printf("\r\nliberando recursos...\r\n")

		// mostra o cursor novamente
		fmt.Print("\033[?25h")

		fmt.Printf("have a nice day!\r\n")
		os.Exit(0)
	}()

	// esconde o cursor
	fmt.Print("\033[?25l")

	timer := time.Tick(time.Duration(300) * time.Millisecond)
	//◐◓◑◒
	//▏▎▍▋
	s := []rune(`◐◓◑◒`)
	slen := len(s)
	i := 0
	for {
		<-timer
		fmt.Print("\r")
		fmt.Print(string(rune(s[i])))
		i++
		if i == slen {
			i = 0
		}
	}
}
