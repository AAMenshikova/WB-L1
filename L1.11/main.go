package main

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{2, 3, 1, 1, 9}
	m := make(map[int]bool)
	for _, valueA := range a {
		m[valueA] = true
	}
	res := []int{}
	for _, valueB := range b {
		if m[valueB] {
			res = append(res, valueB)
			m[valueB] = false
		}
	}
	for _, v := range res {
		println(v)
	}
}