package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCRUDFlow(t *testing.T) {
	s := NewStore()

	createReq := httptest.NewRequest(http.MethodPost, "/items", bytes.NewBufferString(`{"name":"book"}`))
	createRes := httptest.NewRecorder()
	s.Handler(createRes, createReq)
	if createRes.Code != http.StatusCreated {
		t.Fatalf("create status=%d", createRes.Code)
	}

	listReq := httptest.NewRequest(http.MethodGet, "/items", nil)
	listRes := httptest.NewRecorder()
	s.Handler(listRes, listReq)
	if listRes.Code != http.StatusOK {
		t.Fatalf("list status=%d", listRes.Code)
	}

	delReq := httptest.NewRequest(http.MethodDelete, "/items/1", nil)
	delRes := httptest.NewRecorder()
	s.Handler(delRes, delReq)
	if delRes.Code != http.StatusNoContent {
		t.Fatalf("delete status=%d", delRes.Code)
	}
}
