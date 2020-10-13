package main

import "fmt"

func MergeSort(a []int) {
	aux := make([]int, len(a))
	copy(aux, a)
	mergeInto(a, aux)
}

func mergeInto(dst, src []int) {
	if len(dst) < 2 {
		return
	}
	m := len(dst) / 2
	mergeInto(src[:m], dst[:m])
	mergeInto(src[m:], dst[m:])
	j, k := 0, m
	for i := 0; i < len(dst); i++ {
		if j < m && k < len(src) {
			if src[j] < src[k] {
				dst[i], j = src[j], j+1
			} else {
				dst[i], k = src[k], k+1
			}
		} else if j < m {
			dst[i], j = src[j], j+1
		} else {
			dst[i], k = src[k], k+1
		}
	}
}

func main() {
	a := []int{4, 5, 7, 3, 1}
	MergeSort(a)
	fmt.Println(a)
}
