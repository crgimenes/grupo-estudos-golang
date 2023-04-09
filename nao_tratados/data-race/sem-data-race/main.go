package main

/*
teste com o comando
go run -race main.go
*/

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
