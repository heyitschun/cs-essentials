package stringcount

import "testing"

var scTestCases = []struct {
	strings []string
	expect  int
}{
	{
		strings: []string{"ab", "c", "def", "ghij"},
		expect:  10,
	},
}

func TestStringCount(t *testing.T) {
	for _, tc := range scTestCases {
		ans := StringCount(tc.strings, 0)
		if ans != tc.expect {
			t.Errorf("StringCount(%v) = %v; expected %v",
				tc.strings,
				ans,
				tc.expect,
			)
		}
	}
}
