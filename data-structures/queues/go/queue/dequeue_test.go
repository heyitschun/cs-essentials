package queue

import (
	"reflect"
	"testing"
)

// Stack.Pop should return first elements from slice and remove it from
// the stack.
type popTest struct {
	arr    []int
	newArr []int
	popped int
	ok     bool
}

var popTests = []popTest{
	{[]int{1, 2, 3, 4}, []int{2, 3, 4}, 1, true},
	{[]int{-5, 10, 7, -3, 8}, []int{10, 7, -3, 8}, -5, true},
	{[]int{}, []int{}, 0, false},
}

func TestPop(t *testing.T) {
	for _, pq := range popTests {
		s := Queue{pq.arr}
		ans, ok := s.Dequeue()
		if ans != pq.popped || ok != pq.ok {
			t.Error(
				"For", pq.arr,
				"expected Pop() ->", pq.popped,
				"got", ans,
			)
		}
		equal := reflect.DeepEqual(s.elements, pq.newArr)
		if !equal {
			t.Error(
				"For", pq.arr,
				"expected Pop() ->", pq.newArr,
				"got", s.elements,
			)
		}
	}
}
