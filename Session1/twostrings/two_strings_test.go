package twostrings

import "testing"

/*
implement a function that takes two strings s1 and s2 and returns:
- 1 if they are equal.
- 2 if s2 is prefix of s1.
- 3 if s2 is suffix of s1.
- 4 if s1 contains s2.
- -1 if either of them is empty or s2 is longer than s1.
- 0 otherwise
*/

type testCase struct {
	input1 string
	input2 string
	output int
}

func TestString(t *testing.T) {
	testCases := []testCase{
		{"", "", Invlaid},
		{"hello", "", Invlaid},
		{"", "hello", Invlaid},
		{"hell", "hello", Invlaid},
		// ----------
		{"hello", "hello", Equal},
		{"hello", "he", Prefix},
		{"hello", "lo", Suffix},
		{"hello", "el", Substring},
		{"hello", "e", Substring},
		{"hello", "oops", None},
		// ----------
		{"بوتکمپ", "بوتکمپ", Equal},
		{"بوتکمپ", "بوت", Prefix},
		{"بوتکمپ", "کمپ", Suffix},
		{"بوتکمپ", "کم", Substring},
		{"bootcamp", "بوتکمپ", None},
		// ----------
		{"😁🤠", "🤠", Suffix},
	}

	for _, tc := range testCases {
		s1 := tc.input1
		s2 := tc.input2
		expected := tc.output

		if actual := Process(s1, s2); actual != expected {
			t.Errorf("Process(%q,%q) = %d, expected %d", s1, s2, actual, expected)
		}
	}
}
