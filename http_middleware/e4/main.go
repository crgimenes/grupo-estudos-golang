package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type key int

const (
	dataKey key = iota
)

type data struct {
	ValueA string `json:"value_a"`
	ValueB int    `json:"value_b"`
}

func setContextData(r *http.Request, d *data) (ro *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, dataKey, d)
	ro = r.WithContext(ctx)
	return
}

func getContextData(r *http.Request) (d data) {
	d = *r.Context().Value(dataKey).(*data)
	return
}

func middleware1() negroni.Handler {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		d := data{
			ValueA: "valor A",
			ValueB: 42,
		}
		r = setContextData(r, &d)
		next(w, r)
	})
}

func middleware2() negroni.Handler {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		d := getContextData(r)
		d.ValueA += "A"
		r = setContextData(r, &d)
		next(w, r)
	})
}

func middleware3() negroni.Handler {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		d := getContextData(r)
		d.ValueA += "A"
		r = setContextData(r, &d)
		next(w, r)
	})
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	d := getContextData(r)
	j, err := json.MarshalIndent(d, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(j)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	n := negroni.Classic()
	n.Use(middleware1())
	n.Use(middleware2())
	n.Use(middleware3())

	r := mux.NewRouter().StrictSlash(true)
	n.UseHandler(r)

	r.HandleFunc("/", handleMain).Methods("GET")

	fmt.Println("main listen at :8080")
	err := http.ListenAndServe(":8080", n)
	if err != nil {
		fmt.Println(err)
	}
}
