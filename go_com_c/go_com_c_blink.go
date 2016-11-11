package main

/*
    #cgo LDFLAGS: -lmraa
    #include <mraa.h>

    mraa_gpio_context led;

    void config(void) {
        led = mraa_gpio_init(13);
        mraa_gpio_dir(led, MRAA_GPIO_OUT);
    }

    void blink(int b) {
        mraa_gpio_write(led,b);
    }

*/
import "C"

import (
	"fmt"
	"time"
)

func main() {
	C.config()
	for {
		fmt.Print("blink!\r\n")
		C.blink(1)
		time.Sleep(300 * time.Millisecond)
		C.blink(0)
		time.Sleep(300 * time.Millisecond)
	}
}
