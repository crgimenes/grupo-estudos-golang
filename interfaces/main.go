package main

import (
	"fmt"
	"io"
	"strings"
)

type Speaker interface {
	Speak() string
}

type Greeter struct {
	Name string
}

func (g Greeter) Speak() string {
	return "hello " + g.Name
}

func describe(v any) string {
	switch t := v.(type) {
	case fmt.Stringer:
		return t.String()
	case error:
		return t.Error()
	case string:
		return t
	default:
		return "unknown"
	}
}

func readAll(r io.Reader) (string, error) {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, r)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func main() {
	var s Speaker = Greeter{Name: "gopher"}
	v, _ := readAll(strings.NewReader("io.Reader example"))
	fmt.Printf("speak=%q describe=%q io=%q\n", s.Speak(), describe("ok"), v)
}
