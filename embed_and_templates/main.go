package main

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	texttemplate "text/template"
)

//go:embed assets/*.txt
var assets embed.FS

func RenderText(name string) (string, error) {
	tpl, err := texttemplate.New("msg").Parse("Hello, {{.}}")
	if err != nil {
		return "", err
	}

	var b bytes.Buffer
	err = tpl.Execute(&b, name)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

func RenderPlainLink(label string, url string) (string, error) {
	tpl, err := texttemplate.New("link").Parse(`<a href="{{.URL}}">{{.Label}}</a>`)
	if err != nil {
		return "", err
	}

	data := struct {
		Label string
		URL   string
	}{
		Label: label,
		URL:   url,
	}

	var b bytes.Buffer
	err = tpl.Execute(&b, data)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

func RenderHTML(name string) (string, error) {
	tpl, err := template.New("page").Parse("<p>{{.}}</p>")
	if err != nil {
		return "", err
	}

	var b bytes.Buffer
	err = tpl.Execute(&b, name)
	if err != nil {
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
	t, err := RenderText("gopher")
	if err != nil {
		panic(err)
	}

	h, err := RenderHTML("<script>alert(1)</script>")
	if err != nil {
		panic(err)
	}

	fmt.Println(t)
	fmt.Println(h)
}
