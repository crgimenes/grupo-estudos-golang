package main

import (
	"fmt"
	"plugin"
)

func pluginLoader(file string, name string) (s plugin.Symbol, err error) {
	var p *plugin.Plugin
	p, err = plugin.Open(file)
	if err != nil {
		return
	}

	s, err = p.Lookup(name)
	if err != nil {
		return
	}
	return
}

func main() {
	fmt.Println("Hello, main!")

	s, err := pluginLoader("./plugin/plugin.so", "Hello")
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
