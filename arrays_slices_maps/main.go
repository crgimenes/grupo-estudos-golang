package main

import "fmt"

func copySlice(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}

func wordFreq(words []string) map[string]int {
	freq := make(map[string]int)
	for _, w := range words {
		freq[w]++
	}
	return freq
}

func main() {
	arr := [3]int{1, 2, 3}
	slice := arr[:2]
	clone := copySlice(slice)
	freq := wordFreq([]string{"go", "go", "lang"})
	fmt.Printf("arr=%v slice=%v clone=%v go=%d\n", arr, slice, clone, freq["go"])
}
