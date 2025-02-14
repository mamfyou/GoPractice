package binarysearch

import "testing"

/*
implement binary-search for a []string
*/

type testCase struct {
	inputArr       []string
	inputKey       string
	expectedOutput int
}

func TestBinarySearch(t *testing.T) {
	testCases := []testCase{
		{
			inputArr:       nil,
			inputKey:       "",
			expectedOutput: -1,
		},
		{
			inputArr:       []string{},
			inputKey:       "",
			expectedOutput: -1,
		},
		{
			inputArr:       []string{"1"},
			inputKey:       "0",
			expectedOutput: -1,
		},
		{
			inputArr:       []string{"1"},
			inputKey:       "2",
			expectedOutput: -1,
		},
		{
			inputArr:       []string{"1"},
			inputKey:       "1",
			expectedOutput: 0,
		},
		{
			inputArr:       []string{"A", "B"},
			inputKey:       "1",
			expectedOutput: -1,
		},
		{
			inputArr:       []string{"A", "B"},
			inputKey:       "A",
			expectedOutput: 0,
		},
		{
			inputArr:       []string{"A", "B"},
			inputKey:       "B",
			expectedOutput: 1,
		},
		{
			inputArr:       []string{"A", "B"},
			inputKey:       "C",
			expectedOutput: -1,
		},
		{
			inputArr:       []string{"ABC", "BAC", "BCA"},
			inputKey:       "AA",
			expectedOutput: -1,
		},
		{
			inputArr:       []string{"ABC", "BAC", "BCA"},
			inputKey:       "ABC",
			expectedOutput: 0,
		},
		{
			inputArr:       []string{"ABC", "BAC", "BCA"},
			inputKey:       "ABB",
			expectedOutput: -1,
		},
		{
			inputArr:       []string{"ABC", "BAC", "BCA"},
			inputKey:       "BAC",
			expectedOutput: 1,
		},
		{
			inputArr:       []string{"ABC", "BAC", "BCA"},
			inputKey:       "BAD",
			expectedOutput: -1,
		},
		{
			inputArr:       []string{"ABC", "BAC", "BCA"},
			inputKey:       "BCA",
			expectedOutput: 2,
		},
		{
			inputArr:       []string{"ABC", "BAC", "BCA"},
			inputKey:       "BCC",
			expectedOutput: -1,
		},
		{
			inputArr:       []string{"apple", "orange", "pear", "pineapple"},
			inputKey:       "123",
			expectedOutput: -1,
		},
		{
			inputArr:       []string{"apple", "orange", "pear", "pineapple"},
			inputKey:       "apple",
			expectedOutput: 0,
		},
		{
			inputArr:       []string{"apple", "orange", "pear", "pineapple"},
			inputKey:       "pineapple",
			expectedOutput: 3,
		},
		{
			inputArr:       []string{"apple", "orange", "pear", "pineapple"},
			inputKey:       "z",
			expectedOutput: -1,
		},
		{
			inputArr:       []string{"apple", "orange", "pear", "pineapple"},
			inputKey:       "orangz",
			expectedOutput: -1,
		},
		{
			inputArr:       []string{"apple", "orange", "pear", "pineapple", "اکبر", "علی"},
			inputKey:       "apple",
			expectedOutput: 0,
		},
		{
			inputArr:       []string{"apple", "orange", "pear", "pineapple", "اکبر", "علی"},
			inputKey:       "pear",
			expectedOutput: 2,
		},
		{
			inputArr:       []string{"apple", "orange", "pear", "pineapple", "اکبر", "علی"},
			inputKey:       "اکبر",
			expectedOutput: 4,
		},
	}

	for _, tc := range testCases {
		arr := tc.inputArr
		key := tc.inputKey
		expected := tc.expectedOutput

		if actual := BinarySearch(arr, key); actual != expected {
			t.Errorf("BinarySearch(%v,%s) = %d, expected %d", arr, key, actual, expected)
		}
	}
}
