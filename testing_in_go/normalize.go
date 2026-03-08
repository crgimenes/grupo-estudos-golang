package main

import "strings"

func NormalizeSpace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
