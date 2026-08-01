package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCRUDFlow(t *testing.T) {
	store := NewStore()

	createRes := request(t, store, http.MethodPost, "/items", `{"name":"book"}`)
	if createRes.Code != http.StatusCreated {
		t.Fatalf("create status=%d", createRes.Code)
	}

	var created Item
	err := json.NewDecoder(createRes.Body).Decode(&created)
	if err != nil {
		t.Fatal(err)
	}
	if created.ID != 1 || created.Name != "book" {
		t.Fatalf("created item=%+v", created)
	}

	getRes := request(t, store, http.MethodGet, "/items/1", "")
	if getRes.Code != http.StatusOK {
		t.Fatalf("get status=%d", getRes.Code)
	}

	updateRes := request(t, store, http.MethodPut, "/items/1", `{"name":"notebook"}`)
	if updateRes.Code != http.StatusOK {
		t.Fatalf("update status=%d", updateRes.Code)
	}

	listRes := request(t, store, http.MethodGet, "/items", "")
	if listRes.Code != http.StatusOK {
		t.Fatalf("list status=%d", listRes.Code)
	}
	if listRes.Body.String() != "[{\"id\":1,\"name\":\"notebook\"}]\n" {
		t.Fatalf("list body=%q", listRes.Body.String())
	}

	deleteRes := request(t, store, http.MethodDelete, "/items/1", "")
	if deleteRes.Code != http.StatusNoContent {
		t.Fatalf("delete status=%d", deleteRes.Code)
	}

	missingRes := request(t, store, http.MethodGet, "/items/1", "")
	if missingRes.Code != http.StatusNotFound {
		t.Fatalf("missing status=%d", missingRes.Code)
	}
}

func TestRejectsInvalidPayload(t *testing.T) {
	store := NewStore()
	res := request(t, store, http.MethodPost, "/items", `{"name":"   "}`)
	if res.Code != http.StatusBadRequest {
		t.Fatalf("status=%d", res.Code)
	}
}

func request(t *testing.T, store *Store, method string, path string, body string) *httptest.ResponseRecorder {
	t.Helper()

	req := httptest.NewRequest(method, path, bytes.NewBufferString(body))
	res := httptest.NewRecorder()
	store.Handler(res, req)
	return res
}
