all:
	gcc -c -Wall -O2 -lmraa button.c
	ar rvs button.a button.o
	go build button.go

clean:
	rm -f button.o button.a button
