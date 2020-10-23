package main

import "fmt"

// Returns index of target value in array
func BinarySearch(arr []int, target int) int {
	for l, h := 0, len(arr)-1; l <= h; {
		m := (l + h) / 2
		switch {
		case target == arr[m]:
			return m
		case target > arr[m]:
			l = m + 1
		case target < arr[m]:
			h = m - 1
		}
	}

	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := 1

	ans := BinarySearch(arr, t)

	fmt.Println(ans)
}
