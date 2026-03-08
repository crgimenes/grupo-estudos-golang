package main

import "fmt"

func increment(v *int) {
	*v++
}

func parse(flag bool) (int, error) {
	result := 10
	if flag {
		result := 20 // shadowing intencional para exemplo
		_ = result
	}
	return result, nil
}

func main() {
	value := 1
	increment(&value)
	parsed, _ := parse(true)
	fmt.Printf("value=%d parsed=%d\n", value, parsed)
}
