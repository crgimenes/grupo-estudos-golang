package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	var out bytes.Buffer
	err := run([]string{"-name", "go"}, strings.NewReader(""), &out)
	if err != nil {
		t.Fatal(err)
	}
	if out.String() != "hello, go\n" {
		t.Fatalf("out=%q", out.String())
	}
}
