package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func f(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("exit by Ctrl + C")
			return
		case <-ticker.C:
			fmt.Println(i)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	go f(ctx, &wg)
	wg.Wait()
}