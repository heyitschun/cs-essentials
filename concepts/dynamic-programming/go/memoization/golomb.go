package memoization

// Golomb calcutates the Nth number from the Golomb sequence
func Golomb(n int, g map[int]int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	currGolomb := g[n]

	if currGolomb != 0 {
		return currGolomb
	} else {
		newGolomb := 1 + Golomb(n-Golomb(Golomb(n-1, g), g), g)
		g[n] = newGolomb
		return newGolomb
	}
}
