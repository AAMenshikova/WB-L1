package main

import (
	"sync"
)

func main() {
	var (
		wg sync.WaitGroup
		mu sync.Mutex
	)
	wg.Add(1000)
	m := make(map[int]int)
	for i := range 1000 {
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}(i)
	}
	wg.Wait()
}