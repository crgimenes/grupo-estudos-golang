package main

import (
	"fmt"
	"time"

	"github.com/kr/beanstalk"
)

func main() {
	c, err := beanstalk.Dial("tcp", "127.0.0.1:11300")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		err = c.Close()
		if err != nil {
			fmt.Println()
		}
	}()

	id, err := c.Put([]byte("Olá mundo!"), 1, 0, time.Duration(120)*time.Second)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(id)
}
