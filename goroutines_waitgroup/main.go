package main

import "sync"

func ConcurrentSquare(values []int) []int {
	out := make([]int, len(values))
	var wg sync.WaitGroup
	wg.Add(len(values))
	for i, v := range values {
		go func(idx, val int) {
			defer wg.Done()
			out[idx] = val * val
		}(i, v)
	}
	wg.Wait()
	return out
}
