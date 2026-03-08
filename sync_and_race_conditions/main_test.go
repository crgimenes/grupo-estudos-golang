package main

import (
	"sync"
	"testing"
)

func TestCounterConcurrent(t *testing.T) {
	var c Counter
	var wg sync.WaitGroup

	workers := 100
	wg.Add(workers)
	for range workers {
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}
	wg.Wait()

	if got := c.Value(); got != workers {
		t.Fatalf("counter = %d, want %d", got, workers)
	}
}

func TestFillSyncMap(t *testing.T) {
	m := FillSyncMap([]string{"go", "lang"})
	v, ok := m.Load("go")
	if !ok || v.(int) != 2 {
		t.Fatalf("unexpected map value: %#v", v)
	}
}
