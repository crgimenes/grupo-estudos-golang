package main

import "fmt"

func classify(n int) string {
	if n < 0 {
		return "negative"
	}
	switch {
	case n == 0:
		return "zero"
	case n%2 == 0:
		return "even"
	default:
		return "odd"
	}
}

func sum(values []int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

func main() {
	fmt.Printf("classify(4)=%s sum=%d bit=%d\n", classify(4), sum([]int{1, 2, 3}), 1<<3)
}
