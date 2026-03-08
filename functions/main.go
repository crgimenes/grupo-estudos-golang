package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func sum(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

func makeAdder(base int) func(int) int {
	return func(v int) int {
		return base + v
	}
}

func init() {
	_ = "init runs before main"
}

func main() {
	q, err := divide(10, 2)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	add10 := makeAdder(10)
	fmt.Printf("divide=%.1f sum=%d closure=%d\n", q, sum(1, 2, 3), add10(5))
}
