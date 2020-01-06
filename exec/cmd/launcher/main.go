package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"
)

var mutex sync.Mutex

func write(w io.Writer, msg string) {
	mutex.Lock()
	w.Write([]byte(msg))
	mutex.Unlock()
}

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command(filepath.Join(pwd, "helper"))
	cmd.Stderr = os.Stderr
	stdin, err := cmd.StdinPipe()
	if nil != err {
		log.Fatalf("Error obtaining stdin: %s", err.Error())
	}
	stdout, err := cmd.StdoutPipe()
	if nil != err {
		log.Fatalf("Error obtaining stdout: %s", err.Error())
	}
	reader := bufio.NewReader(stdout)

	go func() {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			s := scanner.Text()
			log.Printf("helper: %q\n", s)
			switch s {
			case "read":
				write(stdin, "ok\n")
			case "ping":
				write(stdin, "pong\n")
			}
		}
	}()

	go func() {
		<-time.After(5 * time.Second)
		write(stdin, "quit\n")
	}()

	if err := cmd.Start(); nil != err {
		log.Fatalf("Error starting program: %s, %s", cmd.Path, err.Error())
	}
	cmd.Wait()
}
