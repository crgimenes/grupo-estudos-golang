package main

/*
#cgo LDFLAGS: -lmraa ras.a
#include "ras.h"
*/
import "C"
import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	const maxAngle = float32(180.0)
	rasValue := float32(0.0)

	keeprunning := true

	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, os.Interrupt)
		// espera pelo sinal
		<-sc
		keeprunning = false
	}()

	C.config()

	for keeprunning {
		rasValue = float32(C.readRAS())
		fmt.Printf("Valor: %f angulo: %f\r\n",
			rasValue,
			rasValue*maxAngle/1023.0)

		time.Sleep(50000 * time.Microsecond)
	}

	C.removeRAS()

}
