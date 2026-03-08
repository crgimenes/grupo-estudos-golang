package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func UserService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"service":"user","status":"ok"}`))
}

func BillingService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"service":"billing","status":"ok"}`))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/user/health", UserService)
	mux.HandleFunc("/billing/health", BillingService)

	srv := httptest.NewServer(mux)
	defer srv.Close()

	for _, path := range []string{"/user/health", "/billing/health"} {
		resp, err := http.Get(srv.URL + path)
		if err != nil {
			fmt.Println("request error:", err)
			continue
		}
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Printf("%s -> %d %s\n", path, resp.StatusCode, string(body))
	}
}
