package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("UNIX box?\r\n")
	switch os := runtime.GOOS; os {
	case "darwin":
		fallthrough
	case "freebsd":
		fallthrough
	case "openbsd":
		fallthrough
	case "plan9":
		fmt.Printf("YES!\r\n.")
	case "linux":
		fmt.Printf("almost...\r\n")
	default:
		fmt.Printf("not at all...\r\n")
	}
}
