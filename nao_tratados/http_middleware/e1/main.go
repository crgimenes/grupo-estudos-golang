package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleMain(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("{\"value\":42}\n"))
	if err != nil {
		fmt.Println("error handleMain", err)
	}
}

func handleHealthcheck(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("{\"status\":\"ok\"}\n"))
	if err != nil {
		fmt.Println("error handleHealthcheck", err)
	}
}

func applicationJSON(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	}
}

func basicAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/healthcheck" {
			h.ServeHTTP(w, r)
			return
		}

		user, pass, ok := r.BasicAuth()
		if !ok || user != "admin" || pass != "admin" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, `{"error": "Unauthorized"}`)
			return
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		h.ServeHTTP(w, r)
	}
}

func main() {

	http.HandleFunc("/", applicationJSON(basicAuth(handleMain)))
	http.HandleFunc("/healthcheck", applicationJSON(handleHealthcheck))

	fmt.Println("main listen at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
