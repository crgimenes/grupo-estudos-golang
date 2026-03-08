package main

import (
	"bytes"
	"html/template"
	"net/http"
)

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Request-ID", "demo-id")
		next.ServeHTTP(w, r)
	})
}

func SessionHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "sid", Value: "abc", HttpOnly: true, Path: "/"})
	_, _ = w.Write([]byte("ok"))
}

func RenderHome(name string) (string, error) {
	tpl, err := template.New("home").Parse("<h1>Olá, {{.}}</h1>")
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	if err := tpl.Execute(&b, name); err != nil {
		return "", err
	}
	return b.String(), nil
}
