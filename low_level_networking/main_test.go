package main

import (
	"io"
	"net"
	"testing"
)

func TestHandleLine(t *testing.T) {
	server, client := net.Pipe()
	defer client.Close()

	errCh := make(chan error, 1)
	go func() { errCh <- HandleLine(server) }()

	if _, err := client.Write([]byte("go\n")); err != nil {
		t.Fatal(err)
	}
	respBytes, err := io.ReadAll(client)
	if err != nil {
		t.Fatal(err)
	}
	resp := string(respBytes)
	if resp != "echo:go" {
		t.Fatalf("resp = %q", resp)
	}
	if err := <-errCh; err != nil {
		t.Fatal(err)
	}
}
