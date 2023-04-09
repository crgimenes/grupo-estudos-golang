package main

import "fmt"

var a = 1

func ret10() (int, error) {
	return 10, nil
}

func main() {
	x := 1
	if a == 1 {
		x, err := ret10()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("x:", x)
	}
	fmt.Println("x:", x)
}
