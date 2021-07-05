package trinum

import "testing"

var tnTestCases = []struct {
	n      int
	expect int
}{
	{
		n:      6,
		expect: 21,
	},
	{
		n:      7,
		expect: 28,
	},
}

func TestTriangularNumbers(t *testing.T) {
	for _, tc := range tnTestCases {
		ans := TriangularNumbers(tc.n)
		if ans != tc.expect {
			t.Errorf("TriangularNumbers(%v) returns %v; expected %v",
				tc.n,
				ans,
				tc.expect,
			)
		}
	}
}
