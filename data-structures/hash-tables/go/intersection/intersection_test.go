package intersection

import (
	"reflect"
	"testing"
)

type test struct {
	a   []int
	b   []int
	ans []int
}

var tests = []test{
	{
		[]int{1, 2, 3, 4, 5},
		[]int{0, 2, 4, 6, 8},
		[]int{2, 4},
	},
	{
		[]int{1, 2, 3, 4, 5},
		[]int{1, 2, 3, 4, 5},
		[]int{1, 2, 3, 4, 5},
	},
	{
		[]int{-10, -1, 5, 10, 400, 1000},
		[]int{10},
		[]int{10},
	},
	{
		[]int{1, 2, 3, 4, 5},
		[]int{6, 7, 8, 9, 10},
		[]int{},
	},
}

func TestIntersection(t *testing.T) {
	for _, q := range tests {
		ans := Intersection(q.a, q.b)
		equal := reflect.DeepEqual(ans, q.ans)
		if !equal {
			t.Error(
				"For", q.a,
				"and", q.b,
				"expected", q.ans,
				"got", ans,
			)
		}
	}
}
