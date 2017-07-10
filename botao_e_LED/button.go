package main

/*
#cgo LDFLAGS: -lmraa button.a
#include "button.h"
*/
import "C"
import "fmt"

func main() {
	C.config()

	for {
		if C.readButton() == 1 {
			C.writeLED(1)
			fmt.Println("Botão pressionado.")
		} else {
			C.writeLED(0)
			fmt.Println("Botão liberado.")
		}
	}
}
