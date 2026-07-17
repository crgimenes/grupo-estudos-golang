package main_test

import (
	"context"
	"os/exec"
	"testing"
	"time"
)

func TestProgramOutput(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "go", "run", ".")
	got, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("go run . failed: %v\n%s", err, got)
	}

	want := "{\n" +
		"  \"id\": 7,\n" +
		"  \"name\": \"Ada\",\n" +
		"  \"email\": \"ada@example.com\"\n" +
		"}\n" +
		"<user>\n" +
		"  <id>7</id>\n" +
		"  <name>Ada</name>\n" +
		"  <email>ada@example.com</email>\n" +
		"</user>\n"
	if string(got) != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
