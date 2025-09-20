package main

import (
	"fmt"
)

func main() {
	nums := [5]int{2, 4, 6, 8, 10}
	ch := make(chan int)
	for _, n := range nums {
		go func(x int) {
			ch <- x * x
		}(n)
	}
	for range nums {
		fmt.Println(<-ch)
	}
}