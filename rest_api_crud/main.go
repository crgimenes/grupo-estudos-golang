package main

import (
	"encoding/json"
	"net/http"
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
		for _, it := range s.items {
			list = append(list, it)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(list)
	case http.MethodPost:
		var in Item
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil || strings.TrimSpace(in.Name) == "" {
			http.Error(w, "invalid payload", http.StatusBadRequest)
			return
		}
		in.ID = s.next
		s.next++
		s.items[in.ID] = in
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(in)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Store) handleItem(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/items/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if r.Method == http.MethodDelete {
		if _, ok := s.items[id]; !ok {
			http.NotFound(w, r)
			return
		}
		delete(s.items, id)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}
