package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadFirstLine(r io.Reader) (string, error) {
	s := bufio.NewScanner(r)
	if !s.Scan() {
		if err := s.Err(); err != nil {
			return "", err
		}
		return "", io.EOF
	}
	return s.Text(), nil
}

func main() {
	line, err := ReadFirstLine(strings.NewReader("hello\nworld\n"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stdout, line)
}
