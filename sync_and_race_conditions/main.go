package main

import "sync"

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
