package main

func ShortestPath(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}

	return ShortestPath(m-1, n) + ShortestPath(m, n-1)
}
