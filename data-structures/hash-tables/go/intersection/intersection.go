package intersection

// Intersection returns and array that contains all values present in
// input arrays `a` and `b`.
// For the sake of simplicity, assume that `a` and `b` and always sorted in
// ascending order.
// Returns `[]int`
func Intersection(a []int, b []int) []int {
	// take largest array
	var bigArr []int
	var smallArr []int
	answer := []int{}
	table := make(map[int]bool)

	if len(a) >= len(b) {
		bigArr = a
		smallArr = b
	} else {
		bigArr = b
		smallArr = a
	}

	for _, i := range bigArr {
		table[i] = true
	}

	for _, j := range smallArr {
		if table[j] {
			answer = append(answer, j)
		}
	}

	return answer
}
