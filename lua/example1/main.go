package main

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

func main() {
	l := lua.NewState()
	defer l.Close()
	err := l.DoString(`
		ola="ola mundo"
		print(ola)
		ola="ola golang"
		`)
	if err != nil {
		fmt.Println(err)
		return
	}
	ola := l.GetGlobal("ola")
	fmt.Println(ola.String())
}
