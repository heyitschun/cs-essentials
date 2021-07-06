package insertion

import (
	"fmt"
	"reflect"
	"testing"
)

type test struct {
	array  []int
	answer []int
}

var tests = []test{
	{[]int{4, 1, 7, 2, 3}, []int{1, 2, 3, 4, 7}},
	{[]int{5, 14, 3, 9, 7, 7}, []int{3, 5, 7, 7, 9, 14}},
	{[]int{0, 1, -1}, []int{-1, 0, 1}},
}

func TestInsertionSort(t *testing.T) {
	for _, pair := range tests {
		ans := InsertionSort(pair.array)
		fmt.Println(pair.array, "->", ans)
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
