package roman

import "testing"

/*
implement a function that converts a number to its Roman form.
https://leetcode.com/problems/integer-to-roman/description
*/

type testCase struct {
	input  int
	output string
}

func TestToRoman(t *testing.T) {
	testCases := []testCase{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{9, "IX"},
		{10, "X"},
		{40, "XL"},
		{45, "XLV"},
		{50, "L"},
		{58, "LVIII"},
		{90, "XC"},
		{101, "CI"},
		{450, "CDL"},
		{506, "DVI"},
		{900, "CM"},
		{1994, "MCMXCIV"},
	}

	for _, tc := range testCases {
		n := tc.input
		expected := tc.output

		if actual := ToRoman(n); actual != expected {
			t.Errorf("ToRoman(%d) = %q, expected %q", n, actual, expected)
		}
	}
}
