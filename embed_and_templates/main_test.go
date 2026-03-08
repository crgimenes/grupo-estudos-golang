package main

import (
	"strings"
	"testing"
)

func TestRenderText(t *testing.T) {
	got, err := RenderText("go")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "Hello, go" {
		t.Fatalf("got %q", got)
	}
}

func TestRenderHTMLAutoEscape(t *testing.T) {
	got, err := RenderHTML("<b>x</b>")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(got, "&lt;b&gt;x&lt;/b&gt;") {
		t.Fatalf("expected escaped output, got %q", got)
	}
}
