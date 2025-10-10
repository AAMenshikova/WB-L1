package main

import (
	"fmt"
	"sync"
	"time"
)

func f(stop chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; ; i++ {
		select {
		case <-stop:
			fmt.Println("exit by channel")
			return
		default: 
			fmt.Println(i)
		}
		time.Sleep(time.Second)
	}
}

func main() {
	stop := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	go f(stop, &wg)
	time.Sleep(time.Second * 3)
	close(stop)
	wg.Wait()
}