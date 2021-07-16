package memoization

import "testing"

type golombTest struct {
	n   int
	ans int
}

var golombTests = []golombTest{
	{4, 3},
	{6, 4},
	{9, 5},
}

func TestGolomb(t *testing.T) {
	for _, gt := range golombTests {
		golombs := make(map[int]int)
		ans := Golomb(gt.n, golombs)
		if ans != gt.ans {
			t.Error(
				"For", gt.n,
				"expected Golomb ->", gt.ans,
				"got", ans,
			)
		}
	}
}
