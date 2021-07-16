package quicksort

// SortableSlice is an ADT that contains a slice of ints.
// It behaves like a normal slice, except that it has a built-in sorting
// method.
type SortableSlice struct {
	elements []int
}

// `partition` takes the slice of ints and partitions them for sorting.
func (s *SortableSlice) partition(lp int, rp int) int {
	// pivot index is always set to the right point for consistency
	pivotIndex := rp
	pivotValue := s.elements[pivotIndex]

	// move right pointer one position to the left
	rp--

	for {
		for s.elements[lp] < pivotValue {
			lp++
		}
		for s.elements[rp] > pivotValue {
			rp--
			if rp < 0 {
				break
			}
		}

		if lp >= rp {
			break
		} else {
			s.elements[lp], s.elements[rp] = s.elements[rp], s.elements[lp]
			lp++
		}
	}

	s.elements[lp], s.elements[pivotIndex] = s.elements[pivotIndex], s.elements[lp]

	return lp
}

// Quicksort will call `partition` on subsections of the array,
// so it cannot be assumed that the left and right pointers are always
// at both far ends of the array. For this reason, they have to be method
// arguments.
func (s *SortableSlice) Quicksort(lp int, rp int) {
	if rp-lp <= 0 {
		return
	}

	// old left pointer becomes new pivot point
	pivotIndex := s.partition(lp, rp)

	// Recursively call Quicksort
	s.Quicksort(lp, pivotIndex-1)
	s.Quicksort(pivotIndex+1, rp)
}
