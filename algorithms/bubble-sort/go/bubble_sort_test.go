package bubblesort

import (
	"reflect"
	"testing"
)

type testpair struct {
	array  []int
	answer []int
}

var tests = []testpair{
	{[]int{4, 1, 7, 3, 2}, []int{1, 2, 3, 4, 7}},
	{[]int{500, 3, 23, 98, 48, 35}, []int{3, 23, 35, 48, 98, 500}},
}

func TestBubbleSort(t *testing.T) {
	for _, pair := range tests {
		ans := BubbleSort(pair.array)
		equal := reflect.DeepEqual(ans, pair.answer)
		if !equal {
			t.Error(
				"For", pair.array,
				"expected", pair.answer,
				"got", ans,
			)
		}
	}
}
