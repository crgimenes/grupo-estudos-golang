package main

import "testing"

func TestDecodeJSON(t *testing.T) {
	u, err := DecodeJSON([]byte(`{"id":1,"name":"go"}`))
	if err != nil {
		t.Fatal(err)
	}
	if u.ID != 1 || u.Name != "go" {
		t.Fatalf("unexpected user: %#v", u)
	}
}

func TestEncodeXML(t *testing.T) {
	x, err := EncodeXML(User{ID: 1, Name: "go"})
	if err != nil {
		t.Fatal(err)
	}
	if len(x) == 0 {
		t.Fatal("empty xml")
	}
}
