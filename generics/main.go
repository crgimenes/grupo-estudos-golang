package main

import (
	"cmp"
	"fmt"
)

func Min[T cmp.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func MapSlice[T any, U any](in []T, fn func(T) U) []U {
	out := make([]U, len(in))
	for i, v := range in {
		out[i] = fn(v)
	}
	return out
}

func main() {
	nums := []int{2, 4, 6}
	asText := MapSlice(nums, func(v int) string {
		return fmt.Sprintf("n=%d", v)
	})

	fmt.Printf("min-int=%d min-string=%q mapped=%v\n", Min(10, 3), Min("go", "golang"), asText)
}
