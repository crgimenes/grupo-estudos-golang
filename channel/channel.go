package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
	a principal função dos channels é manter uma comunicação entre as goroutines

	declarando channel com buffer
	c1 := make(chan string, 10)

	declarando channel sem buffer
	c2 := make(chan string)

	enviando msg para o channel
	c1 <- "Arroz"
	c2 <- "Feijão"

	recebendo dados de um canal
	data1 := <-c1
	data2 := <-c2
*/

var goRoutine sync.WaitGroup

func main() {
	rand.Seed(time.Now().Unix())

	//Criando um channel com buffer para se comunicar com as goroutines
	projetos := make(chan string, 10)

	//Cinco goroutines são criadas
	goRoutine.Add(5)
	for i := 1; i <= 5; i++ {
		go funcionarios(projetos, i)
	}

	//Channel é alimentado com os valores para manter a comunicação
	for j := 1; j <= 10; j++ {
		projetos <- fmt.Sprintf("Projeto :%d", j)
	}

	//Fecha o canal para que as goroutines sejam finalizadas
	close(projetos)
	goRoutine.Wait()
}

func funcionarios(projetos chan string, funcionarios int) {
	defer goRoutine.Done()
	for {
		//Aguarda até que receba mensagens no channel
		projeto, result := <-projetos

		if result == false {
			//Caso o channel esteja vazio ou seja fechado
			fmt.Printf("Funcionário : %d : Saiu\n", funcionarios)
			return
		}

		fmt.Printf("Funcionário : %d : Iniciou   %s\n", funcionarios, projeto)

		sleep := rand.Int63n(50)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Println("Tempo de espera", sleep, "ms")
		fmt.Println("")

		fmt.Printf("Funcionário : %d : Completou %s\n", funcionarios, projeto)
	}
}
