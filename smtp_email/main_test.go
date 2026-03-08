package main

import (
	"net/smtp"
	"strings"
	"testing"
)

func TestBuildMessage(t *testing.T) {
	msg := string(BuildMessage("a@x", "b@y", "hi", "body"))
	if !strings.Contains(msg, "Subject: hi") {
		t.Fatalf("unexpected message: %q", msg)
	}
}

func TestSMTPAuth(t *testing.T) {
	var _ smtp.Auth = SMTPAuth("u", "p", "smtp.example.com")
}
