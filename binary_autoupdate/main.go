package main

import "strings"

func IsNewerVersion(current, latest string) bool {
	return strings.TrimSpace(current) != strings.TrimSpace(latest)
}
