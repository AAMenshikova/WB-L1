package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func f(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("exit by context cancel")
			return
		default:
			fmt.Println(i)
		}
		time.Sleep(time.Second)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(1)
	go f(ctx, &wg)
	time.Sleep(time.Second * 3)
	cancel()
	wg.Wait()
}