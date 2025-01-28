package main

import (
	lua "github.com/yuin/gopher-lua"
)

func concat(l *lua.LState) int {
	a := l.ToString(1)
	b := l.ToString(2)
	ret := a + b
	l.Push(lua.LString(ret))
	return 1
}

func sum(l *lua.LState) int {
	a := l.ToInt(1)
	b := l.ToInt(2)
	ret := a + b
	l.Push(lua.LNumber(ret))
	return 1
}

func main() {
	l := lua.NewState()
	defer l.Close()

	l.SetGlobal("concat", l.NewFunction(concat))
	l.SetGlobal("sum", l.NewFunction(sum))

	err := l.DoString(`

		print("concat(\"test1\",\"test2\"): " .. concat("test1","test2"))
		print("sum(2,2): " .. sum(2,2))

	`)
	if err != nil {
		panic(err)
	}
}
