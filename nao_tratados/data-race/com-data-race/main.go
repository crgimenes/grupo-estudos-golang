package main

/*
teste com o comando
go run -race main.go
*/

func main() {
	var x int

	go func() {
		x++
	}()

	x++
}
