package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func ReadEnv(key string) (string, bool) {
	v, ok := os.LookupEnv(key)
	return v, ok
}

func RunEcho(msg string) (string, error) {
	cmd := exec.Command("echo", msg)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return out.String(), nil
}

func main() {
	os.Setenv("APP_MODE", "dev")
	v, _ := ReadEnv("APP_MODE")
	out, _ := RunEcho(v)
	fmt.Print(out)
}
