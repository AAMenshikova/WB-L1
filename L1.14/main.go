package main

import "fmt"

func recognize(input interface{}) {
	switch input.(type) {
	case int:
		fmt.Print("Type is int\n")
	case string:
		fmt.Print("Type is string\n")
	case bool:
		fmt.Print("Type is bool\n")
	case chan int:
		fmt.Print("Type is chan int\n")
	case chan bool:
		fmt.Print("Type is chan bool\n")
	case chan interface{}:
		fmt.Print("Type is chan interface{}\n")
	default:
		fmt.Print("Unknown type\n")
	}
}

func main() {
	recognize(5)
	recognize("hello")
	recognize(true)
	recognize(make(chan int))
	recognize(make(chan bool))
	recognize(make(chan interface{}))
	recognize(3.14)
}