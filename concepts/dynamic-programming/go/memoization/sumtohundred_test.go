package memoization

import "testing"

type sumHundredTest struct {
	arr []int
	ans int
}

var sumHundredTests = []sumHundredTest{
	{
		[]int{49, 50, 51},
		99,
	},
	{
		[]int{200, 201},
		0,
	},
	{
		[]int{10, 30, 40, 70},
		80,
	},
}

// SumToHundred sums and array of integers, as long as a particular number
// does not bring it above 100 (that number is ignored).
func TestSumToHundred(t *testing.T) {
	t.Skip()
	for _, sht := range sumHundredTests {
		ans := SumToHundred(sht.arr)
		println(ans)
		if ans != sht.ans {
			t.Error(
				"For", sht.arr,
				"expected SumToHundred ->", sht.ans,
				"got", ans,
			)
		}
	}
}
