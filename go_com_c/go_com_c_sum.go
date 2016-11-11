package main

/*
int sum(int a,int b) {
	return a+b;
}
*/
import "C"

import (
	"fmt"
)

func main() {
	r := C.sum(2, 2)
	fmt.Printf("sum(2,2) = %v\r\n", r)

}
