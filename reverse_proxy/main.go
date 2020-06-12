package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		w.Header().Set("X-MEUHEADER", "olá mundo!")
		p.ServeHTTP(w, r)
	}
}

func main() {
	remote, err := url.Parse("http://localhost:8080")
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	http.HandleFunc("/", handler(proxy))

	fmt.Println("iniciando proxy na porta 80")
	err = http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err)
	}
}
