package main

import "sync"
import "time"


func main() {

	/*
		Exemplo de goroutines controladas através WaitGroup. 
		WaitGroup aguarda a execução de goroutines até que elas sejam finalizadas
	*/

	println("Inicio")

	//Cria um grupo de espera para aguardar o processamento de todos as goroutines
	var waitGroup sync.WaitGroup
	
	//Define o numero de goroutines que deve ser aguardado.
	waitGroup.Add(2) /*Mude para 3 para aguardar as 3 goroutines ou mude para 0 para não aguardar nenhuma delas*/
	
	
	go func() {
		time.Sleep(2 * time.Second)
		println(" -> foo")
		//Decrementa o contador do WaitGroup.
		waitGroup.Done()
	}()

	go func() {
		time.Sleep(3 * time.Second)
		println(" -> bar")
		waitGroup.Done()
	}()
	
	go func() {
		time.Sleep(1 * time.Second)
		println(" -> qux")
		waitGroup.Done()
	}()
	
	//Segura a execução até que o contador do WaitGroup chegue a 0.
	waitGroup.Wait()
	
	println("Fim")
	
}
