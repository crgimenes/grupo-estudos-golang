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

	id, body, err := c.Reserve(5 * time.Second)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(id, string(body))

	err = c.Delete(id)
	if err != nil {
		fmt.Println(err)
	}
}
