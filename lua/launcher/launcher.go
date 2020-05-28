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
	waitForList = make(map[string]chan struct{})
	helperQuit  = make(chan struct{})
	mutex       sync.Mutex
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
			for k, c := range waitForList {
				if strings.Contains(s, k) {
					c <- struct{}{}
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

func infiniteLoop(l *lua.LState) int {
	c := make(chan struct{})
	<-c
	return 0
}

func waitFor(l *lua.LState) int {
	a := l.ToString(1)

	mutex.Lock()
	c := make(chan struct{})
	waitForList[a] = c
	mutex.Unlock()

	<-c

	mutex.Lock()
	delete(waitForList, a)
	mutex.Unlock()

	res := lua.LString(a)
	l.Push(res)
	return 1
}

func trigger(l *lua.LState) int {
	a := l.ToString(1)
	f := l.ToFunction(2)

	waitForList[a] = make(chan struct{})

	go func() {
		for {
			<-waitForList[a]
			err := l.CallByParam(lua.P{
				Fn:      f,    // Lua function
				NRet:    0,    // number of returned values
				Protect: true, // return err or panic
			})
			if err != nil {
				fmt.Println(err)
			}
		}
	}()

	res := lua.LString(a)
	l.Push(res)
	return 1
}

func runLua() error {
	l := lua.NewState()
	defer l.Close()

	l.SetGlobal("trigger", l.NewFunction(trigger))
	l.SetGlobal("waitFor", l.NewFunction(waitFor))
	l.SetGlobal("infiniteLoop", l.NewFunction(infiniteLoop))

	return l.DoFile("script.lua")
}

func main() {

	go func() {
		err := execHelper()
		if err != nil {
			log.Println(err)
		}
	}()

	err := runLua()
	if err != nil {
		log.Println(err)
	}

	c := make(chan struct{})
	<-c

}
