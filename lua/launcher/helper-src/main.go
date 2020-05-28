package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for {
		if i == 0 {
			fmt.Println("it is a test!")
		} else {
			fmt.Println("loop", i)
			if i > 9 {
				i = 0
				continue
			}
			if i == 5 {
				fmt.Println("golang + lua rocks!")
			}
		}
		i++
		<-time.After(1 * time.Second)
	}
}
