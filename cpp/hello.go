package hello

/*
#include "helloc.h"
*/
import "C"

import "unsafe"

type hello struct {
	helloPtr C.helloPtr
}

func New() (r hello) {
	r.helloPtr = C.helloInit()
	return
}

func (h *hello) Free() {
	C.helloFree(unsafe.Pointer(h.helloPtr))
}

func (h *hello) Print() {
	C.helloPrint(unsafe.Pointer(h.helloPtr))
}
