package binarysearch

import (
	"testing"
)

type testpair struct {
	array  []int
	target int
	answer int
}

var tests = []testpair{
	{[]int{1, 3, 5, 6}, 5, 2},
	{[]int{100, 101, 102, 103, 104, 105}, 7, -1},
	{[]int{-10, -5, 0, 5, 10}, -5, 1},
}

func TestBinarySearch(t *testing.T) {
	for _, pair := range tests {
		ans, _ := BinarySearch(pair.array, pair.target)
		if ans != pair.answer {
			t.Error(
				"For", pair.array,
				"expected", pair.answer,
				"got", ans,
			)
		}
	}
}
