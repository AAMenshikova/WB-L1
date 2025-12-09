package main

import (
	"strings"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	runes := []rune(v)
	if len(runes) > 100 {
		runes = runes[:100]
	}
	justString = string(runes)
}

func createHugeString(n int) string {
	return strings.Repeat("a", n)
}

func main() {
	someFunc()
}