package stack

import "testing"

type readTest struct {
	s   []int
	ans int
	ok  bool
}

var readTests = []readTest{
	{[]int{1, 2, 3, 4}, 4, true},
	{[]int{-5, 8, 10, 3, -6}, -6, true},
	{[]int{}, 0, false},
}

func TestRead(t *testing.T) {
	for _, rt := range readTests {
		s := Stack{rt.s}
		ans, ok := s.Read()
		if ans != rt.ans || ok != rt.ok {
			t.Error(
				"For", rt.s,
				"expected Read() ->", rt.ans,
				"got", ans,
			)
		}
	}
}
