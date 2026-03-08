package main

import (
	"bufio"
	"fmt"
	"regexp"
	"slices"
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

func main() {
	data := "app=go-study\nmode=dev\ninvalid-line\nworkers=4\n"
	parsed := ParseKeyValues(data)
	keys := make([]string, 0, len(parsed))
	for k := range parsed {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	for _, k := range keys {
		fmt.Printf("%s=%s\n", k, parsed[k])
	}
}
