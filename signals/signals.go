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

		// Aqui você fecha o banco de dados, libera memória, etc...

		fmt.Printf("have a nice day!\r\n")
		os.Exit(0)
	}()

	fmt.Printf("Pressione ^C para terminar\r\n")

	for {
		/*
			fica colocando pontos na tela a cada segundo só
			para mostrar que o programa esta rodando
		*/
		time.Sleep(1 * time.Second)
		fmt.Printf(".")
	}

}
