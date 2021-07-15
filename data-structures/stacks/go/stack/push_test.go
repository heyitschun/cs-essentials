package stack

import (
	"reflect"
	"testing"
)

// Stack.Push should add an element to the top of the stack
type pushTest struct {
	arr    []int
	newArr []int
	newVal int
}

var pushTests = []pushTest{
	{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4, 5}, 5},
	{[]int{-5, 10, 7, -3, 8}, []int{-5, 10, 7, -3, 8, -6}, -6},
}

func TestPush(t *testing.T) {
	for _, pq := range pushTests {
		s := Stack{pq.arr}
		s.Push(pq.newVal)
		equal := reflect.DeepEqual(s.elements, pq.newArr)
		if !equal {
			t.Error(
				"For", pq.arr,
				"pushing", pq.newVal,
				"expected", pq.newArr,
				"got", s.elements,
			)
		}
	}
}
