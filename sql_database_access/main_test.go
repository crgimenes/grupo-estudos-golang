package main

import "testing"

func TestInsertAndList(t *testing.T) {
	s := NewUserStore()
	s.Insert("a")
	s.Insert("b")
	users := s.List()
	if len(users) != 2 {
		t.Fatalf("len = %d", len(users))
	}
	if users[0].Name != "a" || users[1].Name != "b" {
		t.Fatalf("unexpected users: %#v", users)
	}
}
