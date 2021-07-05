package filter

import "testing"

var feTestCases = []struct {
	array  []int
	expect []int
}{
	{
		array:  []int{1, 2, 3, 4},
		expect: []int{2, 4},
	},
}

func TestFilterEvens(t *testing.T) {
	for _, tc := range feTestCases {
		ans := FilterEvens(tc.array)

		if ans != tc.expect {
			t.Errorf("FilterEvens(%v) returns %v; expected %v",
				tc.array,
				ans,
				tc.expect,
			)
		}
	}
}
