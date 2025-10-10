package main

import (
	"fmt"
	"sync"
)

func f(wg *sync.WaitGroup) {
	defer func(){ if r := recover(); r != nil { fmt.Println("handled panic:", r) } }()
	defer wg.Done()
	for i := 0; ; i++ {
		if i == 3 {
			panic("exit by panic")
		}
		fmt.Println(i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go f(&wg)
	wg.Wait()
}