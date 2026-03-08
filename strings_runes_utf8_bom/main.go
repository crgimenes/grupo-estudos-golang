package main

import (
	"fmt"
	"unicode/utf8"
)

func runeCount(s string) int {
	return utf8.RuneCountInString(s)
}

func hasUTF8BOM(b []byte) bool {
	if len(b) < 3 {
		return false
	}
	return b[0] == 0xEF && b[1] == 0xBB && b[2] == 0xBF
}

func main() {
	text := "Go é simples"
	fmt.Printf("bytes=%d runes=%d\n", len(text), runeCount(text))
	for i, r := range text {
		fmt.Printf("%d:%q ", i, r)
	}
	fmt.Println()
}
