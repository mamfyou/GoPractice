package bubblesort

import (
	"slices"
	"testing"
)

/*
implement ascending in-place bubble-sort for a []string
*/

type testCase struct {
	input  []string
	output []string
}

func TestBubbleSort(t *testing.T) {
	testCases := []testCase{
		{
			input:  nil,
			output: nil,
		},
		{
			input:  []string{},
			output: []string{},
		},
		{
			input:  []string{"1"},
			output: []string{"1"},
		},
		{
			input:  []string{"A", "B"},
			output: []string{"A", "B"},
		},
		{
			input:  []string{"B", "A"},
			output: []string{"A", "B"},
		},
		{
			input:  []string{"ABC", "BAC", "BCA"},
			output: []string{"ABC", "BAC", "BCA"},
		},
		{
			input:  []string{"ABC", "BCA", "BAC"},
			output: []string{"ABC", "BAC", "BCA"},
		},
		{
			input:  []string{"BAC", "ABC", "BCA"},
			output: []string{"ABC", "BAC", "BCA"},
		},
		{
			input:  []string{"BAC", "BCA", "ABC"},
			output: []string{"ABC", "BAC", "BCA"},
		},
		{
			input:  []string{"BCA", "ABC", "BAC"},
			output: []string{"ABC", "BAC", "BCA"},
		},
		{
			input:  []string{"BCA", "BAC", "ABC"},
			output: []string{"ABC", "BAC", "BCA"},
		},
		{
			input:  []string{"pineapple", "apple", "pear", "orange"},
			output: []string{"apple", "orange", "pear", "pineapple"},
		},
		{
			input:  []string{"orange", "pear", "apple", "pineapple"},
			output: []string{"apple", "orange", "pear", "pineapple"},
		},
		{
			input:  []string{"علی", "اکبر", "اکبر", "علی", "اکبر"},
			output: []string{"اکبر", "اکبر", "اکبر", "علی", "علی"},
		},
		{
			input:  []string{"علی", "orange", "اکبر", "pear", "apple", "pineapple", "apple", "pear", "اکبر"},
			output: []string{"apple", "apple", "orange", "pear", "pear", "pineapple", "اکبر", "اکبر", "علی"},
		},
	}

	for _, tc := range testCases {
		input := tc.input
		expected := tc.output

		actual := make([]string, len(input))
		copy(actual, input)

		Sort(actual)

		if !slices.Equal(actual, expected) {
			t.Errorf("Sort(%v) = %v, expected %v", input, actual, expected)
		}
	}
}
