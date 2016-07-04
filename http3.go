package main

import (
	"flag"
	"log"
	"net/http"
)

func setHSTSHeader(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")

		h.ServeHTTP(w, r)
	}
}

func main() {

	var port = flag.String("port", "9999", "Define what TCP port to bind to")
	var root = flag.String("root", "assets/", "Define the root filesystem path")

	flag.Parse()

	assets := setHSTSHeader(http.FileServer(http.Dir(*root)))
	http.Handle("/", assets)

	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
