package main

import (
	"fmt"
	"strconv"
	"strings"
)

func IsNewerVersion(current, latest string) bool {
	cmp, err := CompareVersions(current, latest)
	return err == nil && cmp < 0
}

func CompareVersions(a, b string) (int, error) {
	va, err := parseVersion(a)
	if err != nil {
		return 0, err
	}
	vb, err := parseVersion(b)
	if err != nil {
		return 0, err
	}
	for i := range 3 {
		if va[i] < vb[i] {
			return -1, nil
		}
		if va[i] > vb[i] {
			return 1, nil
		}
	}
	return 0, nil
}

func parseVersion(v string) ([3]int, error) {
	var out [3]int
	parts := strings.Split(strings.TrimSpace(v), ".")
	if len(parts) != 3 {
		return out, fmt.Errorf("invalid version format: %q", v)
	}
	for i := range out {
		n, err := strconv.Atoi(parts[i])
		if err != nil {
			return out, fmt.Errorf("invalid version value %q: %w", parts[i], err)
		}
		out[i] = n
	}
	return out, nil
}

func main() {
	current := "1.4.2"
	latest := "1.5.0"
	fmt.Printf("current=%s latest=%s update_required=%v\n", current, latest, IsNewerVersion(current, latest))
}
