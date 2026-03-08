package main

import "testing"

func TestStructFieldNames(t *testing.T) {
	got, err := StructFieldNames(User{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 2 || got[0] != "ID" || got[1] != "Name" {
		t.Fatalf("unexpected names: %#v", got)
	}
}

func TestStructFieldNamesInvalid(t *testing.T) {
	if _, err := StructFieldNames(10); err == nil {
		t.Fatal("expected error for non-struct")
	}
}
