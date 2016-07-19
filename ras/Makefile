all:
	gcc -c -Wall -O2 -lmraa ras.c
	ar rvs ras.a ras.o
	go build ras.go

clean:
	rm -f ras.o ras.a pwm

