package bubblesort

// BubbleSort sorts a slice of integers into ascending order.
// Returns the sorted slice.
func BubbleSort(a []int) []int {
	unsortedIndex := len(a) - 1
	sorted := false

	for !sorted {
		sorted = true
		for i := 0; i < unsortedIndex; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				sorted = false
			}
		}
		unsortedIndex--
	}

	return a
}
