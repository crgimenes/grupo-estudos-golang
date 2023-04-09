package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf("t é uma variavel tipo %T\r\n", t)
	fmt.Println((t.Format(time.RFC3339)))

	// só pode ser piada! "01/02 03:04:05PM '06 -0700" sério?

	fmt.Println(t.Format("2006-01-02 15:04:05-0700"))
}
