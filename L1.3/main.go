package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func worker(id int, jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d: %s\n", id, j)
	}
	fmt.Printf("Worker %d: chanel closed, stopping\n", id)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Count of arguments must be 1")
		os.Exit(1)
	}
	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numWorkers <= 0 {
		fmt.Println("Count of workers must be positive integer")
		os.Exit(1)
	}

	jobs := make(chan string)
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		go worker(i, jobs, &wg)
	}
	fmt.Printf("Started %d workers\n", numWorkers)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	fmt.Println("Press Ctrl+C to stop")

	cnt := 1
	for {
		select {
		case <-sigCh:
			close(jobs)
			wg.Wait()
			fmt.Println("All workers done!")
			return
		case jobs <- fmt.Sprintf("Data item %d", cnt):
			cnt++
			time.Sleep(500 * time.Millisecond)
		}

	}
}
