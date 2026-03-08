package main

import "fmt"

type User struct {
	ID   int
	Name string
}

type UserStore struct {
	nextID int
	data   map[int]User
}

func NewUserStore() *UserStore {
	return &UserStore{nextID: 1, data: map[int]User{}}
}

func (s *UserStore) Insert(name string) User {
	u := User{ID: s.nextID, Name: name}
	s.nextID++
	s.data[u.ID] = u
	return u
}

func (s *UserStore) List() []User {
	out := make([]User, 0, len(s.data))
	for i := 1; i < s.nextID; i++ {
		if u, ok := s.data[i]; ok {
			out = append(out, u)
		}
	}
	return out
}

func main() {
	store := NewUserStore()
	store.Insert("gopher")
	fmt.Println(len(store.List()))
}
