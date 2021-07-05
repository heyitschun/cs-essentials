package factorial

import (
	"testing"
)

type testpair struct {
	n      int
	answer int
}

var tests = []testpair{
	{0, 1},
	{2, 2},
	{5, 120},
	{10, 3628800},
}

func TestFactorial(t *testing.T) {
	for _, pair := range tests {
		f := Factorial(pair.n)
		if f != pair.answer {
			t.Error(
				"For", pair.n,
				"expected", pair.answer,
				"got", f,
			)
		}
	}
}
