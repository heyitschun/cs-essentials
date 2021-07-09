package alphabet

import "testing"

type test struct {
	s string
	a rune
}

var tests = []test{
	{"the quick brown box jumps over a lazy dog", 'f'},
}

func TestAlphabet(t *testing.T) {
	for _, q := range tests {
		ans := Alphabet(q.s)
		if ans != q.a {
			t.Error(
				"For", q.s,
				"expected", q.a,
				"got", ans,
			)
		}
	}
}
