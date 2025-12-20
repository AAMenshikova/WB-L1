package main

import "fmt"

func quickSort(a []int) []int {
	n := len(a)
	if n <= 1 {
		out := make([]int, n)
		copy(out, a)
		return out
	}
	pivot := a[n/2]
	less := make([]int, 0, n)
	equal := make([]int, 0, n)
	greater := make([]int, 0, n)
	for _, x := range a {
		switch {
		case x < pivot:
			less = append(less, x)
		case x > pivot:
			greater = append(greater, x)
		default:
			equal = append(equal, x)
		}
	}
	less = quickSort(less)
	greater = quickSort(greater)
	out := make([]int, 0, n)
	out = append(out, less...)
	out = append(out, equal...)
	out = append(out, greater...)
	return out
}

func main() {
	arr := []int{5, 2, 9, 1, 5, 6, 3, 3, 0, -2}
	fmt.Println(quickSort(arr))
	fmt.Println(arr)
}
