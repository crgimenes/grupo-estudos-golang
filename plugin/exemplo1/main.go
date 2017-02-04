package main

import (
	"fmt"
	"plugin"
)

func main() {
	fmt.Println("Hello, main!")

	var p *plugin.Plugin
	var err error
	p, err = plugin.Open("./plugin/plugin.so")
	if err != nil {
		fmt.Println(err)
		return
	}

	var s plugin.Symbol
	s, err = p.Lookup("Hello")
	if err != nil {
		fmt.Println(err)
		return
	}

	Hello := s.(func() error)

	err = Hello()
	if err != nil {
		fmt.Println(err)
		return
	}

}
