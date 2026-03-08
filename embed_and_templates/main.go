package main

import (
	"bytes"
	"embed"
	"fmt"
	template2 "html/template"
	"text/template"
)

//go:embed assets/*.txt
var assets embed.FS

func RenderText(name string) (string, error) {
	tpl, err := template.New("msg").Parse("Hello, {{.}}")
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	if err := tpl.Execute(&b, name); err != nil {
		return "", err
	}
	return b.String(), nil
}

func RenderHTML(name string) (string, error) {
	tpl, err := template2.New("page").Parse("<p>{{.}}</p>")
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	if err := tpl.Execute(&b, name); err != nil {
		return "", err
	}
	return b.String(), nil
}

func EmbeddedReadme() (string, error) {
	b, err := assets.ReadFile("assets/hello.txt")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func main() {
	t, _ := RenderText("gopher")
	h, _ := RenderHTML("<script>alert(1)</script>")
	fmt.Println(t)
	fmt.Println(h)
}
