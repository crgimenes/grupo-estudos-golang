package main

import "testing"

func TestBirthday(t *testing.T) {
	p := Person{Name: "A", Age: 10}
	p.Birthday()
	if p.Age != 11 {
		t.Fatalf("age = %d, want 11", p.Age)
	}
}

func TestEmbedding(t *testing.T) {
	e := Employee{Person: Person{Name: "A", Age: 20}, Role: "Dev"}
	if got := e.Label(); got != "A(20)" {
		t.Fatalf("label = %q", got)
	}
}
