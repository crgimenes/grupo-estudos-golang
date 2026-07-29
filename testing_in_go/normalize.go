package space

import "strings"

// NormalizeSpace collapses any run of whitespace into one ASCII space.
func NormalizeSpace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
