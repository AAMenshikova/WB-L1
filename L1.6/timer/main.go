package main

import (
	"fmt"
	"sync"
	"time"
)

func f(t int, wg *sync.WaitGroup) {
	defer wg.Done()
	timer := time.NewTimer(time.Duration(t) * time.Second)
	defer timer.Stop()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for i := 0; ; i++ {
		select {
		case <-timer.C:
			fmt.Println("exit by timer")
			return
		case <-ticker.C:
			fmt.Println(i)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go f(3, &wg)
	wg.Wait()
}