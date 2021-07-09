package nonduplicate

import "testing"

type test struct {
	s string
	a rune
}

var tests = []test{
	{"minimum", 'n'},
}

func TestNonduplicates(t *testing.T) {
	for _, q := range tests {
		ans := Nonduplicate(q.s)
		if ans != q.a {
			t.Error(
				"For", q.s,
				"expect", q.a,
				"got", ans,
			)
		}
	}
}
