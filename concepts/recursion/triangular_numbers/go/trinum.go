package trinum

func TriangularNumbers(n int) int {
	if n == 1 {
		return 1
	}
	return TriangularNumbers(n-1) + n
}
