package stack

import (
	"testing"
)

type lintTest struct {
	str string
	err error
}

var lintTests = []lintTest{
	{"(var x = {y: [1, 2, 3]})", nil},
	{"var x = {y: [1, 2, 3]})", SyntaxError2},
	{"(var x = y: [{1, 2, 3]})", SyntaxError3},
	{"(var x = {y: [1, 2, 3]}", SyntaxError1},
}

func TestLinter(t *testing.T) {
	for _, lt := range lintTests {
		err := BraceChecker(lt.str)
		if err != lt.err {
			t.Error(
				"For", lt.str,
				"expected BraceChecker ->", lt.err,
				"got", err,
			)
		}
	}
}
