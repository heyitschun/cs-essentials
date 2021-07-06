package selection

func SelectionSort(a []int) []int {
	var lowestIdxNumber int
	var temp int
	for i := 0; i < len(a); i++ {
		lowestIdxNumber = i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[lowestIdxNumber] {
				lowestIdxNumber = j
			}
		}
		if lowestIdxNumber != i {
			temp = a[i]
			a[i] = a[lowestIdxNumber]
			a[lowestIdxNumber] = temp
		}
	}

	return a
}
