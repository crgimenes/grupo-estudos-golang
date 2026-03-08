package main

import (
	"fmt"
	"strings"
)

func divideOrPanic(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}

func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered panic: %v", r)
		}
	}()
	result = divideOrPanic(a, b)
	return result, nil
}

func deferOrder() string {
	parts := []string{}
	func() {
		defer func() { parts = append(parts, "third") }()
		defer func() { parts = append(parts, "second") }()
		defer func() { parts = append(parts, "first") }()
		parts = append(parts, "start")
	}()
	return strings.Join(parts, ",")
}

func main() {
	_, err := safeDivide(10, 0)
	fmt.Println(err != nil)
}
