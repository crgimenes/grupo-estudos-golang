package main

import (
	"log"
	"net/http"
)

func main() {
	assets := http.FileServer(http.Dir("assets/"))
	http.Handle("/", assets)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
