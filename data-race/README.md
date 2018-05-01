# Data Race

Data race é uma condição que ocorre quando duas threads tentam acessar a mesma area da memória ao mesmo tempo e pode causar vários problemas e resultados imprevisíveis.

Felizmente Go vem completo com um detector de data race para testar seu código.

Veja o exemplo abaixo:

```go
package main

func main() {
	var x int

	go func() {
		x++
	}()

	x++
}

```

Nesse exemplo a variável x é acessada por dois pontos, um pela função main() e outro pela goroutine que criamos e para detectar o data race usamos o seguinte comando:

```sh
go run -race main.go
```

Esse comando vai analizar seu código e exibir um alerta indicando onde é possível acontecer data race, alem do run você pode usar o parâmetro -race também com test, build e install.

Para corrigir o problema existem várias maneiras mas a mais simples e tradicional é usar mutex e cercar com lock e unlock todos os lugares onde acessamos a variável x.

Veja o exemplo do código corrigido:

```go
package main

import "sync"

func main() {
	var x int
	var m sync.Mutex

	go func() {
		m.Lock()
		x++
		m.Unlock()
	}()

	m.Lock()
	x++
	m.Unlock()
}
```
