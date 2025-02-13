package ugly

import "testing"

/*
implement a function that checks if an integer value is ugly!
a number is NOT ugly when it has a prime factor other than {2,3,5}.
*/

type testCase struct {
	input  int
	output bool
}

func TestIsUgly(t *testing.T) {
	testCases := []testCase{
		{0, false},
		{1, false},
		{2, true},
		{3, true},
		{4, true},
		{5, true},
		{6, true},
		{7, false},
		{8, true},
		{9, true},
		{10, true},
		{11, false},
		{12, true},
		{13, false},
		{14, false},
		{15, true},
		{16, true},
		{17, false},
		{18, true},
		{19, false},
		{20, true},
		{21, false},
		{22, false},
		{23, false},
		{24, true},
		{25, true},
		{26, false},
		{27, true},
		{28, false},
		{29, false},
		{30, true},
	}

	for _, tc := range testCases {
		n := tc.input
		expected := tc.output

		if actual := IsUgly(n); actual != expected {
			t.Errorf("IsUgly(%d) = %t, expected %t", n, actual, expected)
		}
	}
}
