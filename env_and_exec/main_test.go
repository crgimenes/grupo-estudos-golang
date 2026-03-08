package main

import (
	"strings"
	"testing"
)

func TestReadEnv(t *testing.T) {
	t.Setenv("X_MODE", "ok")
	v, ok := ReadEnv("X_MODE")
	if !ok || v != "ok" {
		t.Fatalf("got (%q, %v)", v, ok)
	}
}

func TestRunEcho(t *testing.T) {
	got, err := RunEcho("go")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "go") {
		t.Fatalf("got %q", got)
	}
}
