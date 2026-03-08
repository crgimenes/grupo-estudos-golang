package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func run(args []string, in io.Reader, out io.Writer) error {
	fs := flag.NewFlagSet("cli_project", flag.ContinueOnError)
	name := fs.String("name", "", "name to greet")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if strings.TrimSpace(*name) == "" {
		return errors.New("-name is required")
	}
	fmt.Fprintf(out, "hello, %s\n", *name)
	return nil
}

func main() {
	if err := run(os.Args[1:], os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
