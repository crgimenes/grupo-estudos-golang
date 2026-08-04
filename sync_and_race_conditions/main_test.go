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

	got := c.Value()
	if got != workers {
		t.Fatalf("counter = %d, want %d", got, workers)
	}
}

func TestIncrementConcurrently(t *testing.T) {
	got := IncrementConcurrently(100)
	if got != 100 {
		t.Fatalf("count = %d, want 100", got)
	}
}

func TestFillSyncMap(t *testing.T) {
	m := FillSyncMap([]string{"go", "lang"})
	v, ok := m.Load("go")
	if !ok {
		t.Fatal("missing map value")
	}

	got := v.(int)
	if got != 2 {
		t.Fatalf("map value = %d, want 2", got)
	}
}
