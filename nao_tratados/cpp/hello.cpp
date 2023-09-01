#include <iostream>
#include "hello.hpp"
#include "helloc.h"

hello::hello(void) {
	std::cout << "construtor da classe hello" << std::endl;
}

hello::~hello(void) {
    std::cout << "destrutor da classe hello" << std::endl;
}

void hello::Print(void) {
    std::cout << "membro da classe hello" << std::endl;
}

/*
int main(void) {

	hello *h = new hello();
	h->Print();
	delete h;
	return 0;
}
*/
