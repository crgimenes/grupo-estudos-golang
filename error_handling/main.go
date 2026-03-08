package main

import (
	"errors"
	"fmt"
	"strconv"
)

var ErrUserNotFound = errors.New("user not found")

type ValidationError struct {
	Field string
	Msg   string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s: %s", e.Field, e.Msg)
}

func findUserName(id int) (string, error) {
	if id <= 0 {
		return "", ValidationError{Field: "id", Msg: "must be greater than zero"}
	}
	if id != 1 {
		return "", fmt.Errorf("find user %d: %w", id, ErrUserNotFound)
	}
	return "gopher", nil
}

func parseAge(input string) (int, error) {
	age, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("parse age %q: %w", input, err)
	}
	if age < 0 {
		return 0, ValidationError{Field: "age", Msg: "must be non-negative"}
	}
	return age, nil
}

func main() {
	_, err := findUserName(2)
	fmt.Println(errors.Is(err, ErrUserNotFound))
}
