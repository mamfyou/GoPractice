package subset

import "testing"

/*
implement a function that takes two slices s1 and s2 and checks if s1 is subset of s2.
*/

type testCase struct {
	input1 []int
	input2 []int
	output bool
}

func TestIsSubsetOf(t *testing.T) {
	testCases := []testCase{
		{
			input1: nil,
			input2: nil,
			output: true,
		},
		{
			input1: []int{},
			input2: []int{},
			output: true,
		},
		{
			input1: []int{1},
			input2: []int{1},
			output: true,
		},
		{
			input1: []int{1},
			input2: []int{2},
			output: false,
		},
		{
			input1: []int{1},
			input2: []int{},
			output: false,
		},
		{
			input1: []int{1, 2, 3},
			input2: []int{1, 2, 3, 4, 5},
			output: true,
		},
		{
			input1: []int{1, 6},
			input2: []int{1, 2, 3, 4, 5},
			output: false,
		},
		{
			input1: []int{6, 1},
			input2: []int{5, 4, 3, 2, 1},
			output: false,
		},
		{
			input1: []int{},
			input2: []int{1, 2, 3},
			output: true,
		},
		{
			input1: []int{4, 5, 6},
			input2: []int{4, 5, 6},
			output: true,
		},
		{
			input1: []int{1, 1, 2, 2, 3, 3},
			input2: []int{4, 3, 3, 3, 2, 2, 2, 1, 1, 1},
			output: true,
		},
	}

	for _, tc := range testCases {
		s1 := tc.input1
		s2 := tc.input2
		expected := tc.output

		if actual := IsSubsetOf(s1, s2); actual != expected {
			t.Errorf("IsSubsetOf(%v,%v) = %t, expected %t", s1, s2, actual, expected)
		}
	}
}
