package main

import (
	"math/rand"
	"sync"
	"time"
)

func rotina(wg *sync.WaitGroup, x int) {
	// Espera um tempo aleatório
	rt := rand.Int31n(1000)
	time.Sleep(time.Duration(rt) * time.Millisecond)

	println("rotina", x)
	wg.Done()
}

func main() {

	println("Inicio")
	rand.Seed(time.Now().Unix())

	var wg sync.WaitGroup

	wg.Add(3)

	for i := 1; i < 4; i++ {
		go rotina(&wg, i)
	}

	wg.Wait()

	println("Fim")
}
