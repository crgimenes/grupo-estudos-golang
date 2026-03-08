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
	ids := []int{1, 2, 0}
	for _, id := range ids {
		name, err := findUserName(id)
		if err != nil {
			if errors.Is(err, ErrUserNotFound) {
				fmt.Printf("id=%d -> not found\n", id)
				continue
			}
			var ve ValidationError
			if errors.As(err, &ve) {
				fmt.Printf("id=%d -> validation error on %s\n", id, ve.Field)
				continue
			}
			fmt.Printf("id=%d -> unexpected error: %v\n", id, err)
			continue
		}
		fmt.Printf("id=%d -> user=%s\n", id, name)
	}

	for _, input := range []string{"25", "-1", "abc"} {
		age, err := parseAge(input)
		if err != nil {
			fmt.Printf("age=%q -> %v\n", input, err)
			continue
		}
		fmt.Printf("age=%q -> %d\n", input, age)
	}
}
