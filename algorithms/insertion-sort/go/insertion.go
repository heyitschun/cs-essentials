package insertion

// InsertionSort sort an array of integers in ascending order.
// Returns an array of integers.
func InsertionSort(a []int) []int {
	var temp int
	var position int
	for i := 1; i < len(a); i++ {
		temp = a[i]
		position = i - 1
		for position >= 0 {
			if a[position] > temp {
				a[position+1] = a[position]
				position = position - 1
			} else {
				break
			}
		}
		a[position+1] = temp
	}

	return a
}
