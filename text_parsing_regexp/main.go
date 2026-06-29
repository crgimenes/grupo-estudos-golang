package main

import (
	"bufio"
	"fmt"
	"regexp"
	"slices"
	"strings"
)

var keyValuePattern = regexp.MustCompile(`^([a-zA-Z0-9_]+)=([^=]+)$`)

func ParseKeyValues(data string) (map[string]string, error) {
	out := map[string]string{}
	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		match := keyValuePattern.FindStringSubmatch(line)
		if len(match) != 3 {
			continue
		}

		out[match[1]] = match[2]
	}

	err := scanner.Err()
	if err != nil {
		return nil, err
	}

	return out, nil
}

func main() {
	data := "app=go-study\nmode=dev\n# ignored\ninvalid-line\nworkers=4\n"
	parsed, err := ParseKeyValues(data)
	if err != nil {
		panic(err)
	}

	keys := make([]string, 0, len(parsed))
	for key := range parsed {
		keys = append(keys, key)
	}

	slices.Sort(keys)
	for _, key := range keys {
		fmt.Printf("%s=%s\n", key, parsed[key])
	}
}
