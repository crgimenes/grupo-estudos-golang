package main

import (
	"errors"
	"testing"
)

func TestFindUserNameErrorsIs(t *testing.T) {
	_, err := findUserName(2)
	if !errors.Is(err, ErrUserNotFound) {
		t.Fatalf("expected errors.Is(..., ErrUserNotFound) to be true, got err=%v", err)
	}
}

func TestValidationErrorAs(t *testing.T) {
	_, err := parseAge("-1")
	var ve ValidationError
	if !errors.As(err, &ve) {
		t.Fatalf("expected ValidationError, got %v", err)
	}
	if ve.Field != "age" {
		t.Fatalf("field = %q, want %q", ve.Field, "age")
	}
}
