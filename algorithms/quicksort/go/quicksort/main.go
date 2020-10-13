package main

import "fmt"

func Quicksort(a []int) {
	if len(a) < 2 {
		return
	}
	ub := len(a) - 1
	pivot := a[ub]
	i, j := -1, ub
	for i < j {
		for i++; a[i] < pivot; i++ {
		}
		for j--; 0 < j && a[j] > pivot; j-- {
		}
		a[i], a[j] = a[j], a[i]
	}
	a[j], a[i], a[ub] = a[i], pivot, a[j]
	Quicksort(a[:i])
	Quicksort(a[i+1:])
}

func main() {
	a := []int{7, 3, 4, 2, 6, 1, 5, 4}
	Quicksort(a)

	fmt.Println(a)
}
