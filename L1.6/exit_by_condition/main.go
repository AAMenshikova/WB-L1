package main

import (
	"fmt"
	"sync"
	"time"
)

func f(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; ; i++ {
		if i == 3 {
			fmt.Println("exit by condition")
			c <- i
			return
		}
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() {
	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go f(c, &wg)
	<-c
	wg.Wait()
}