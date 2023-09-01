#include "hello.hpp"
#include "helloc.h"

helloPtr helloInit(void) {
    hello *h = new hello();
    return (void*)h;
}

void helloFree(helloPtr h) {
    hello *h_aux = (hello*)h;
    delete h_aux;
}

void helloPrint(helloPtr h) {
    hello *h_aux = (hello*)h;
    h_aux->Print();
}

/*
int main(void) {
    helloPtr h = helloInit();
    helloPrint(h);
    helloFree(h);
    return 0;
}
*/
