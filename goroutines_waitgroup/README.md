# Goroutines com WaitGroup 

Para auxiliar no controle de diversos jobs simultâneos, independentes e ainda com resultados sincronizados podemos lançar mão da WaitGroup.
A WaitGroup aguarda que um determinado número de goroutines sejam finalizadas para prosseguir com a execução. 
O programa principal faz a chamada para .Add() que recebe como parâmetro o número de goroutines que deve aguardar serem finalizadas. 


Exemplo:

```go
waitGroup.Add(3)
```


No exemplo abaixo demostramos a chamada de três goroutines e a chamada para .Done() ao final de cada uma das goroutines. 
.Done() por sua vez realiza o decremento do número de goroutines parametrizados em .Add()
E antes da finalização da execução do programa principal e após a chamada das goroutines, realizamos a chamada para .Wait() que deve fazer com que toda a execução pare neste ponto, aguardando que o decremento gerado por .Done() chegue a 0.

```go
package main

import "sync"
import "time"


func main() {

	println("Inicio")

	//Cria um grupo de espera para aguardar o processamento de todos as goroutines
	var waitGroup sync.WaitGroup
	
	//Define o número de goroutines que deve ser aguardado.
	waitGroup.Add(3) 
	
	
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
```
[Playground](https://play.golang.org/p/z9P_CPTjkG)

---
[Inicio](../README.md)

[< Goroutines](../goroutines/) - [Select >](../select/)
