package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu sync.Mutex
	n  int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	c.n++
	c.mu.Unlock()
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}

func FillSyncMap(values []string) *sync.Map {
	m := &sync.Map{}
	for _, v := range values {
		m.Store(v, len(v))
	}
	return m
}

func IncrementConcurrently(workers int) int {
	var c Counter
	var wg sync.WaitGroup

	wg.Add(workers)
	for range workers {
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}
	wg.Wait()

	return c.Value()
}

func main() {
	count := IncrementConcurrently(100)
	cache := FillSyncMap([]string{"go", "mutex", "race"})

	value, ok := cache.Load("go")
	if !ok {
		panic("missing cache entry")
	}

	fmt.Printf("counter=%d\n", count)
	fmt.Printf("cache[go]=%d\n", value)
}
