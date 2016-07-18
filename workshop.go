package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {

	var port = flag.String("port", "8888", "Define what TCP port to bind to")
	var root = flag.String("root", "assets/", "Define the root filesystem path")

	flag.Parse()

	fmt.Printf("http://localhost:%s/\n", *port)

	http.Handle("/", http.FileServer(http.Dir(*root)))

	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
