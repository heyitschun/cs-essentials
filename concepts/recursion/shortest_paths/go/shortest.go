package main

import "fmt"

func UniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	return UP(m-1, n-1)
}

func UP(m int, n int) int {
	if m < 0 || n < 0 {
		return 0
	} else if m == 0 || n == 0 {
		return 1
	}
	return UP(m-1, n) + UP(m, n-1)
}

func main() {
	p := UniquePaths(3, 7)
	fmt.Println(p)
}
