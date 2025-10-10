package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func f(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; ; i++ {
		if i == 3 {
			fmt.Println("exit by goexit")
			runtime.Goexit()
		}
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go f(&wg)
	wg.Wait()
}