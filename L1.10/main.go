package main

import (
	"fmt"
)

func main() {
	input := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	output := make(map[int][]float64)
	for _, t := range input {
		ind := int(t / 10) * 10
		output[ind] = append(output[ind], t)
	}
	for i := range output {
		fmt.Print(i, ": ", output[i], "\n")
	}
}