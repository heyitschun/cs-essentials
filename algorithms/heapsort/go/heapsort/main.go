package main

import "fmt"

func HeapSort(a []int) {
	if len(a) < 2 {
		return
	}
	maxIndex := len(a) - 1
	for i := (maxIndex - 1) / 2; 0 <= i; i-- {
		trickleDown(a, i, maxIndex)
	}
	for {
		a[0], a[maxIndex] = a[maxIndex], a[0]
		maxIndex--
		if maxIndex <= 0 {
			break
		}
		trickleDown(a, 0, maxIndex)
	}
}

func trickleDown(a []int, i, maxIndex int) {
	tmp := a[i]
	for j := 2*i + 1; j <= maxIndex; j = 2*i + 1 {
		if j < maxIndex && a[j] < a[j+1] {
			j++
		}
		if a[j] <= tmp {
			break
		}
		a[i], i = a[j], j
	}
	a[i] = tmp
}

func main() {
	a := []int{0, 1, 3, 4, 2}

	HeapSort(a)
	fmt.Println(a)
}
