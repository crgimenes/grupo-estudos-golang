package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserService(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	UserService(w, r)
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d", w.Code)
	}
}
