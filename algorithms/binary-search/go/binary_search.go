package binarysearch

import (
	"errors"
	"fmt"
)

// BinarySearch takes an array of integers and searches for a target integer.
// Returns the index of the target (-1 if not found) and an error.
func BinarySearch(a []int, t int) (int, error) {
	lower := 0
	upper := len(a) - 1
	middle := 0

	for lower <= upper {
		fmt.Println(middle)
		middle = (lower + upper) / 2
		if a[middle] < t {
			lower = middle + 1
		} else if a[middle] > t {
			upper = middle - 1
		} else {
			return middle, nil
		}
	}

	return -1, errors.New("Search target not in array.")
}
