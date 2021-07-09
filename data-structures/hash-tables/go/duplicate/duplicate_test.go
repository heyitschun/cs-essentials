package duplicate

import (
	"fmt"
	"testing"
)

type test struct {
	s   []string
	ans string
}

var tests = []test{
	{[]string{"a", "b", "c", "d", "c", "e"}, "c"},
	{[]string{"a", "b", "b", "c", "c"}, "b"},
}

func TestDuplicate(t *testing.T) {
	for _, q := range tests {
		ans := Duplicate(q.s)
		fmt.Printf("%s -> %s\n", q.s, ans)
		if !(ans == q.ans) {
			t.Error(
				"For", q.s,
				"expected", q.ans,
				"got", ans,
			)
		}
	}
}
