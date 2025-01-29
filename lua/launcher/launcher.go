package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
	"sync"

	lua "github.com/yuin/gopher-lua"
)

var (
	triggerList    = make(map[string]*lua.LFunction)
	waitForString  string
	waitForChannel = make(chan struct{})
	helperQuit     = make(chan struct{})
	mutex          sync.Mutex
	l              = lua.NewState()
)

func write(w io.Writer, msg string) {
	mutex.Lock()
	_, err := w.Write([]byte(msg))
	mutex.Unlock()
	if err != nil {
		fmt.Println("write error:", err)
	}
}

func execHelper() error {
	log.Println("starting execHelper")
	cmd := exec.Command("./helper")

	stderr, err := cmd.StderrPipe()
	if nil != err {
		log.Println("error obtaining stderr:", err.Error())
		return err
	}

	stdin, err := cmd.StdinPipe()
	if nil != err {
		log.Println("error obtaining stdin:", err.Error())
		return err
	}

	stdout, err := cmd.StdoutPipe()
	if nil != err {
		log.Println("error obtaining stdout:", err.Error())
		return err
	}
	reader := bufio.NewReader(stdout)
	readerError := bufio.NewReader(stderr)

	go func() {
		scanner := bufio.NewScanner(readerError)
		for scanner.Scan() {
			s := scanner.Text()
			log.Println("helper stderr:", s)
		}
	}()

	go func() {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			s := scanner.Text()
			log.Println("helper stdout:", s)
			switch s {
			case "ready":
				write(stdin, "ok\n")
			case "ping":
				write(stdin, "pong\n")
			}

			mutex.Lock()

			if waitForString != "" &&
				strings.Contains(s, waitForString) {
				waitForChannel <- struct{}{}
			}
			for k, f := range triggerList {
				if strings.Contains(s, k) {
					err := l.CallByParam(lua.P{
						Fn:      f,     // Lua function
						NRet:    0,     // number of returned values
						Protect: false, // return err or panic
					})
					if err != nil {
						fmt.Println(err)
					}
				}
			}
			mutex.Unlock()
		}
	}()

	err = cmd.Start()
	if err != nil {
		e := fmt.Errorf(
			"error starting program: %s, %s",
			cmd.Path,
			err.Error())
		helperQuit <- struct{}{}
		log.Println(e)
		return e
	}
	err = cmd.Wait()
	if err != nil {
		return err
	}
	log.Println("helper exited sending helperQuit message")
	helperQuit <- struct{}{}
	log.Println("exiting execHelper")
	return nil
}

func waitFor(l *lua.LState) int {
	mutex.Lock()
	waitForString = l.ToString(1)
	mutex.Unlock()
	<-waitForChannel
	mutex.Lock()
	res := lua.LString(waitForString)
	waitForString = ""
	mutex.Unlock()
	l.Push(res)
	return 1
}

func trigger(l *lua.LState) int {
	a := l.ToString(1)
	f := l.ToFunction(2)

	mutex.Lock()
	triggerList[a] = f
	mutex.Unlock()

	res := lua.LString(a)
	l.Push(res)
	return 1
}

func runLua() {

	l.SetGlobal("trigger", l.NewFunction(trigger))
	l.SetGlobal("waitFor", l.NewFunction(waitFor))

	err := l.DoFile("script.lua")
	if err != nil {
		fmt.Println(err)
	}

	c := make(chan struct{})
	<-c
}

func main() {

	go func() {
		err := execHelper()
		if err != nil {
			log.Println(err)
		}
	}()

	runLua()

}
