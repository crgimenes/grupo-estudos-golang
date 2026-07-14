package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()
	HealthHandler(w, r)
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d", w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Fatalf("content-type = %q", contentType)
	}

	var body message
	err := json.NewDecoder(w.Body).Decode(&body)
	if err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if body.Status != "ok" {
		t.Fatalf("status body = %q", body.Status)
	}
}

func TestHealthHandlerMethod(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/health", nil)
	w := httptest.NewRecorder()
	HealthHandler(w, r)
	if w.Code != http.StatusMethodNotAllowed {
		t.Fatalf("status = %d", w.Code)
	}
	allow := w.Header().Get("Allow")
	if allow != http.MethodGet {
		t.Fatalf("allow = %q", allow)
	}
}
