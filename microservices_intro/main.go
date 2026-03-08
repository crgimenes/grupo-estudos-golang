package main

import (
	"fmt"
	"net/http"
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
	fmt.Println("microservices intro")
}
