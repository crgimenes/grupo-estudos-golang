package main

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"
)

var (
	//go:embed assets/*
	assets      embed.FS
	assetsFS, _ = fs.Sub(assets, "assets")
)

func printTemplate(page string, w http.ResponseWriter, data any) error {
	tpl, err := template.ParseFS(assets, page)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	return tpl.Execute(w, data)
}

func errorPage(w http.ResponseWriter, msg string) {
	data := struct {
		Message string
	}{
		Message: msg,
	}

	printTemplate("assets/error.html", w, data)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title            string
		SubTitle         string
		Content          string
		Footer           string
		ShowReturnButton bool
		InitPage         bool
		Value            float64
	}{
		Title:            "Olá Mundo!",
		SubTitle:         "Bem-vindo ao Go Web Server",
		Content:          "Este é um exemplo de servidor web em Go",
		Footer:           "Go Web Server",
		ShowReturnButton: false,
		InitPage:         true,
		Value:            math.Pi,
	}

	err := printTemplate("assets/index.html", w, data)
	if err != nil {
		errorPage(w, "Error rendering template")
	}
}

func handlePg2(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title            string
		SubTitle         string
		Content          string
		Footer           string
		ShowReturnButton bool
		InitPage         bool
		Value            float64
	}{
		Title:            "Página 2",
		SubTitle:         "Esta é a página 2",
		Content:          "Esta é a segunda página do servidor web",
		Footer:           "Go Web Server",
		InitPage:         false,
		ShowReturnButton: true,
		Value:            math.E,
	}

	// erro aleatório para demonstrar a página de erro
	if rand.Intn(4) == 1 {
		errorPage(w, "Erro aleatório")
		return
	}

	err := printTemplate("assets/index.html", w, data)
	if err != nil {
		errorPage(w, "Error rendering template")
	}
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /assets/", http.StripPrefix("/assets/",
		http.FileServer(http.FS(assetsFS)),
	))
	mux.Handle("GET /", http.HandlerFunc(handleIndex))
	mux.Handle("GET /pagina2", http.HandlerFunc(handlePg2))

	s := &http.Server{
		Handler:        mux,
		Addr:           ":8080",
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Starting server on %s\n", s.Addr)
	log.Fatal(s.ListenAndServe())
}
