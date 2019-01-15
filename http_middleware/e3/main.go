package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func handleMain(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("{\"value\":42}\n"))
	if err != nil {
		fmt.Println("error handleMain", err)
	}
}

func middleware1() negroni.Handler {
	fmt.Println("carregando middleware 1")
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		fmt.Println("empilhando middleware 1")
		next(w, r)
		fmt.Println("desempilhando middleware 1")
	})
}

func middleware2() negroni.Handler {
	fmt.Println("carregando middleware 2")
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		fmt.Println("empilhando middleware 2")
		next(w, r)
		fmt.Println("desempilhando middleware 2")
	})
}

func middleware3() negroni.Handler {
	fmt.Println("carregando middleware 3")
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		fmt.Println("empilhando middleware 3")
		next(w, r)
		fmt.Println("desempilhando middleware 3")
	})
}

func main() {
	n := negroni.Classic()
	n.Use(middleware1())
	n.Use(middleware2())
	n.Use(middleware3())
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	r := mux.NewRouter().StrictSlash(true)
	n.UseHandler(r)

	r.HandleFunc("/", handleMain).Methods("GET")

	fmt.Println("main listen at :8080")
	err := http.ListenAndServe(":8080", n)
	if err != nil {
		fmt.Println(err)
	}
}
