package bbolt

import (
	"os"
	"testing"
)

func TestBasicStorageFunctions(t *testing.T) {
	defer os.Remove("test.db")
	s, err := New("test.db")
	if err != nil {
		t.Fatalf("New(): %v\n", err)
	}
	err = s.Update("jobs", "testKey", []byte("test 123"))
	if err != nil {
		t.Fatalf("Uodate(): %v\n", err)
	}
	b, err := s.View("jobs", "testKey")
	if err != nil {
		t.Fatalf("View(): %v\n", err)
	}
	if string(b) != "test 123" {
		t.Fatalf("View() expected \"test 123\" but got %q\n", string(b))
	}
	b, err = s.View("test", "testKey")
	if err == nil {
		t.Fatalf("View() expected errors but got nil\n")
	}
	s, err = New("/dev/null")
	if err == nil {
		t.Fatalf("expected error but got nil\n")
	}
}
