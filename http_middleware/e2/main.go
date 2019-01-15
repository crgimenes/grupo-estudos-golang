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

func handleHealthcheck(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("{\"status\":\"ok\"}\n"))
	if err != nil {
		fmt.Println("error handleHealthcheck", err)
	}
}

func applicationJSON() negroni.Handler {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	})
}

func basicAuth() negroni.Handler {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

		if r.URL.Path == "/healthcheck" {
			next(w, r)
			return
		}

		user, pass, ok := r.BasicAuth()
		if !ok || user != "admin" || pass != "admin" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, `{"error": "Unauthorized"}`)
			return
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		next(w, r)
	})
}

func main() {
	n := negroni.Classic()
	n.Use(applicationJSON())
	n.Use(basicAuth())

	r := mux.NewRouter().StrictSlash(true)
	n.UseHandler(r)

	r.HandleFunc("/", handleMain).Methods("GET")
	r.HandleFunc("/healthcheck", handleHealthcheck).Methods("GET")

	fmt.Println("main listen at :8080")
	err := http.ListenAndServe(":8080", n)
	if err != nil {
		fmt.Println(err)
	}
}
