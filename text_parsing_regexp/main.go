package main

import (
	"bufio"
	"regexp"
	"strings"
)

var keyValuePattern = regexp.MustCompile(`^([a-zA-Z0-9_]+)=([^=]+)$`)

func ParseKeyValues(data string) map[string]string {
	out := map[string]string{}
	s := bufio.NewScanner(strings.NewReader(data))
	for s.Scan() {
		m := keyValuePattern.FindStringSubmatch(strings.TrimSpace(s.Text()))
		if len(m) == 3 {
			out[m[1]] = m[2]
		}
	}
	return out
}
