all:
	gcc -c -Wall -O2 -lmraa pwm.c
	ar rvs pwm.a pwm.o
	go build pwm.go

clean:
	rm -f pwm.o pwm.a pwm
