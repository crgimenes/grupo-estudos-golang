package main

import (
	"html/template"
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

func TestRenderPlainLinkDoesNotEscapeByContext(t *testing.T) {
	got, err := RenderPlainLink("docs", "javascript:alert(1)")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := `<a href="javascript:alert(1)">docs</a>`
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
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

func TestRenderHTMLBlocksUnsafeURL(t *testing.T) {
	tpl, err := template.New("link").Parse(`<a href="{{.}}">docs</a>`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var b strings.Builder
	err = tpl.Execute(&b, "javascript:alert(1)")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(b.String(), "#ZgotmplZ") {
		t.Fatalf("expected blocked URL marker, got %q", b.String())
	}
}
