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
		recovered := recover()
		if recovered != nil {
			err = fmt.Errorf("recovered panic: %v", recovered)
		}
	}()

	result = divideOrPanic(a, b)
	return result, nil
}

func deferOrder() string {
	var parts []string
	func() {
		defer func() { parts = append(parts, "third") }()
		defer func() { parts = append(parts, "second") }()
		defer func() { parts = append(parts, "first") }()
		parts = append(parts, "start")
	}()
	return strings.Join(parts, ",")
}

func main() {
	value, err := safeDivide(10, 2)
	if err != nil {
		fmt.Println("unexpected error:", err)
		return
	}
	fmt.Println("10/2 =", value)

	_, err = safeDivide(10, 0)
	if err != nil {
		fmt.Println("recovered:", err)
	}

	fmt.Println("defer order:", deferOrder())
}
