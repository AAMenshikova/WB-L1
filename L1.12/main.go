package main

import "fmt"

func main() {
	input := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool)
	res := []string{}
	for _, i := range input {
		if !set[i] {
			res = append(res, i)
			set[i] = true
		}
	}
	fmt.Print(res)
}