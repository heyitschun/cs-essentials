package quicksort

import (
	"reflect"
	"testing"
)

type quicksortTest struct {
	arr []int
	ans []int
}

var quicksortTests = []quicksortTest{
	{[]int{0, 5, 2, 1, 6, 3}, []int{0, 1, 2, 3, 5, 6}},
	{[]int{50, 7, 16, -1, 7, -30}, []int{-30, -1, 7, 7, 16, 50}},
}

func TestQuicksort(t *testing.T) {
	for _, qt := range quicksortTests {
		s := SortableSlice{qt.arr}
		s.Quicksort(0, len(s.elements)-1)
		ans := s.elements
		equal := reflect.DeepEqual(ans, qt.ans)
		if !equal {
			t.Error(
				"For", qt.arr,
				"expected Quicksort() ->", qt.ans,
				"got", ans,
			)
		}
	}
}
