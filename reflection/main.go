package main

import (
	"fmt"
	"reflect"
)

func StructFieldNames(v any) ([]string, error) {
	t := reflect.TypeOf(v)
	if t == nil {
		return nil, fmt.Errorf("nil type")
	}
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected struct, got %s", t.Kind())
	}
	out := make([]string, 0, t.NumField())
	for field := range t.Fields() {
		out = append(out, field.Name)
	}
	return out, nil
}

type User struct {
	ID   int
	Name string
}

func main() {
	names, _ := StructFieldNames(User{})
	fmt.Println(names)
}
