package main

import (
	"encoding/json"
	"testing"
)

func TestLogRequest(t *testing.T) {
	logger, buf := NewBufferLogger()
	LogRequest(logger, "/health", 200, 3)

	var record map[string]any
	err := json.Unmarshal(buf.Bytes(), &record)
	if err != nil {
		t.Fatal(err)
	}

	checks := map[string]any{
		"level":       "INFO",
		"msg":         "request",
		"path":        "/health",
		"status":      float64(200),
		"duration_ms": float64(3),
	}
	for key, want := range checks {
		got := record[key]
		if got != want {
			t.Fatalf("%s = %v, want %v", key, got, want)
		}
	}

	_, hasTime := record["time"]
	if hasTime {
		t.Fatalf("log contains nondeterministic time field: %s", buf.String())
	}
}
