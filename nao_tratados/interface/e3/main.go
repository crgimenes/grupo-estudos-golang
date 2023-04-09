package main

import (
	"fmt"
	"time"
)

type calc interface {
	Factorial(uint64) uint64
}

type recursive struct{}

func (r recursive) Factorial(n uint64) uint64 {
	if n > 0 {
		return n * r.Factorial(n-1)
	}
	return 1
}

const LIM = 41

type memorization struct {
	facts [LIM]uint64
}

func (m memorization) Factorial(n uint64) uint64 {
	if m.facts[n] != 0 {
		return m.facts[n]
	}
	if n > 0 {
		m.facts[n] = n * m.Factorial(n-1)
		return m.facts[n]
	}
	return 1
}

func printFactorial(c calc, v uint64) {
	start := time.Now()
	fmt.Printf("value %v\n", c.Factorial(v))
	fmt.Printf("%T took %v\n\n", c, time.Since(start))
}

func main() {

	m := memorization{}
	printFactorial(m, 30)

	r := recursive{}
	printFactorial(r, 30)

}
