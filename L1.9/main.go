package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	a := []int{1, 2, 3, 4, 5}
	go func() {
		defer close(ch1)
		for _, c := range a {
			ch1 <- c
		}
	}()
	go func() {
		defer close(ch2)
		for c := range ch1 {
			ch2 <- c * 2
		}
	}()
	for c := range ch2 {
		fmt.Println(c)
	}
}