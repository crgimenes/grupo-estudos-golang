package main

import (
	"fmt"
	"os"

	lua "github.com/yuin/gopher-lua"
)

func main() {

	b, err := os.ReadFile("config.lua")
	if err != nil {
		fmt.Println(err)
		return
	}

	l := lua.NewState()
	defer l.Close()

	err = l.DoString(string(b))
	if err != nil {
		fmt.Println(err)
		return
	}

	ola := l.GetGlobal("ola")
	fmt.Println(ola.String())
}
