package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

/*
inspurado por http://funcoeszz.net/man.html#zzaleatorio
*/

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var min int
	var max int
	var err error
	nargs := len(os.Args)
	switch nargs {
	case 1:
		max = 32767
	case 2:
		max, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
	case 3:
		min, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		max, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println(min + rand.Intn(max-min))
}
