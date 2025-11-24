package main

import "fmt"

func main() {
	var x int
	var y int
	fmt.Scan(&x, &y)
	y = y ^ x
	x = y ^ x
	y = y ^ x
	fmt.Print(x, y)
}