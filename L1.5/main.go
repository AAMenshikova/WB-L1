package main

import (
	"fmt"
	"time"
)

func producer(c chan int, stop <-chan struct{}) {
	for i := 1; ; i++ {
		select {
		case <-stop:
			return
		case c <- i:
			time.Sleep(1 * time.Second)
		}
	}
}

func concumer(c chan int, stop <-chan struct{}) {
	for {
		select {
		case val := <-c:
			fmt.Println(val)
		case <-stop:
			return
		}
	}

}

func main() {
	c := make(chan int)
	stop := make(chan struct{})
	time.AfterFunc(5*time.Second, func() { close(stop) })
	go producer(c, stop)
	concumer(c, stop)
}
