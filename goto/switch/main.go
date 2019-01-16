package main

import (
	"fmt"
)

func main() {
	var i = 2
label:
	switch i {
	case 1:
		fmt.Println("case 1")
	case 2:
		fmt.Println("case 2")
		i = 1
		goto label
	case 3:
		fmt.Println("case 3")
	case 4:
		fmt.Println("case 4")

	}
}
