package main

/*
#cgo LDFLAGS: -lmraa pwm.a
#include "pwm.h"
*/
import "C"

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

var duty float32 = 0.5  // guarda o ciclo de trabalho
var delta float32 = 1.0 // variação do ciclo de trabalho

func main() {

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

		if duty >= 1 {
			duty = 1      // Intencidade do LED
			delta = -0.05 // Diminui a variação do ciclo de trabalho
		} else if duty <= 0 {
			duty = 0      // Intencidade do LED
			delta = +0.05 // Aumenta a variação do ciclo de trabalho
		}

		// Ajusta o ciclo de trabalho
		C.writePWM(C.float(duty))

		time.Sleep(50000 * time.Microsecond)

		duty = duty + delta

		fmt.Printf("ciclo de trabalho: %f\r\n", duty)
	}

	C.writePWM(0)
	C.stopPWM()
}
