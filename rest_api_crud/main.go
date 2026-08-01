package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Store struct {
	mu    sync.Mutex
	next  int
	items map[int]Item
}

func NewStore() *Store {
	return &Store{next: 1, items: map[int]Item{}}
}

func (s *Store) Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/items" {
		s.handleCollection(w, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/items/") {
		s.handleItem(w, r)
		return
	}
	http.NotFound(w, r)
}

func (s *Store) handleCollection(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()

	switch r.Method {
	case http.MethodGet:
		list := make([]Item, 0, len(s.items))
		for _, item := range s.items {
			list = append(list, item)
		}
		sort.Slice(list, func(i, j int) bool {
			return list[i].ID < list[j].ID
		})
		writeJSON(w, http.StatusOK, list)
	case http.MethodPost:
		var input Item
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil || strings.TrimSpace(input.Name) == "" {
			http.Error(w, "invalid payload", http.StatusBadRequest)
			return
		}
		input.ID = s.next
		s.next++
		s.items[input.ID] = input
		writeJSON(w, http.StatusCreated, input)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Store) handleItem(w http.ResponseWriter, r *http.Request) {
	id, ok := itemID(r.URL.Path)
	if !ok {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	item, exists := s.items[id]
	if !exists {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		writeJSON(w, http.StatusOK, item)
	case http.MethodPut:
		var input Item
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil || strings.TrimSpace(input.Name) == "" {
			http.Error(w, "invalid payload", http.StatusBadRequest)
			return
		}
		input.ID = id
		s.items[id] = input
		writeJSON(w, http.StatusOK, input)
	case http.MethodDelete:
		delete(s.items, id)
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func itemID(path string) (int, bool) {
	value := strings.TrimPrefix(path, "/items/")
	id, err := strconv.Atoi(value)
	if err != nil || id < 1 {
		return 0, false
	}
	return id, true
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func main() {
	store := NewStore()

	requests := []struct {
		method string
		path   string
		body   string
	}{
		{method: http.MethodPost, path: "/items", body: `{"name":"notebook"}`},
		{method: http.MethodGet, path: "/items/1"},
		{method: http.MethodPut, path: "/items/1", body: `{"name":"pen"}`},
		{method: http.MethodDelete, path: "/items/1"},
	}

	for _, request := range requests {
		body := bytes.NewBufferString(request.body)
		req := httptest.NewRequest(request.method, request.path, body)
		res := httptest.NewRecorder()
		store.Handler(res, req)

		fmt.Printf("%s %s -> %d\n", request.method, request.path, res.Code)
	}
}
