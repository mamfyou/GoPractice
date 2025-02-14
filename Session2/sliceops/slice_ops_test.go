package sliceops

import (
	"slices"
	"testing"
)

/*
implement functions for the following slice operations: RemoveAll, DropAt and InsertAt
*/

func TestRemoveAll(t *testing.T) {
	testCases := []struct {
		inputArr       []string
		inputKey       string
		expectedOutput []string
	}{
		{
			inputArr:       nil,
			inputKey:       "",
			expectedOutput: nil,
		},
		{
			inputArr:       []string{},
			inputKey:       "Hello",
			expectedOutput: []string{},
		},
		{
			inputArr:       []string{"goodbye"},
			inputKey:       "Hello",
			expectedOutput: []string{"goodbye"},
		},
		{
			inputArr:       []string{"Hello"},
			inputKey:       "Hello",
			expectedOutput: []string{},
		},
		{
			inputArr:       []string{"Hello1", "Hello2"},
			inputKey:       "goodbye",
			expectedOutput: []string{"Hello1", "Hello2"},
		},
		{
			inputArr:       []string{"Hello1", "Hello2"},
			inputKey:       "Hello1",
			expectedOutput: []string{"Hello2"},
		},
		{
			inputArr:       []string{"Hello1", "Hello2"},
			inputKey:       "Hello2",
			expectedOutput: []string{"Hello1"},
		},
		{
			inputArr:       []string{"Hello1", "Hello1"},
			inputKey:       "Hello1",
			expectedOutput: []string{},
		},
		{
			inputArr:       []string{"apple", "orange", "apple", "pear", "pineapple", "apple"},
			inputKey:       "apple",
			expectedOutput: []string{"orange", "pear", "pineapple"},
		},
	}

	for _, tc := range testCases {
		arr := tc.inputArr
		key := tc.inputKey
		expected := tc.expectedOutput

		if actual := RemoveAll(arr, key); !slices.Equal(actual, expected) {
			t.Errorf("RemoveAll(%v,%s) = %v, expected %v", arr, key, actual, expected)
		}
	}
}

func TestDropAt(t *testing.T) {
	testCases := []struct {
		inputArr       []string
		inputIdx       int
		expectedOutput []string
	}{
		{
			inputArr:       []string{"hello"},
			inputIdx:       0,
			expectedOutput: []string{},
		},
		{
			inputArr:       []string{"Hello1", "Hello2"},
			inputIdx:       0,
			expectedOutput: []string{"Hello2"},
		},
		{
			inputArr:       []string{"Hello1", "Hello2"},
			inputIdx:       1,
			expectedOutput: []string{"Hello1"},
		},
		{
			inputArr:       []string{"apple", "orange", "pear", "pineapple"},
			inputIdx:       0,
			expectedOutput: []string{"orange", "pear", "pineapple"},
		},
		{
			inputArr:       []string{"apple", "orange", "pear", "pineapple"},
			inputIdx:       1,
			expectedOutput: []string{"apple", "pear", "pineapple"},
		},
		{
			inputArr:       []string{"apple", "orange", "pear", "pineapple"},
			inputIdx:       2,
			expectedOutput: []string{"apple", "orange", "pineapple"},
		},
		{
			inputArr:       []string{"apple", "orange", "pear", "pineapple"},
			inputIdx:       3,
			expectedOutput: []string{"apple", "orange", "pear"},
		},
	}

	for _, tc := range testCases {
		arr := tc.inputArr
		index := tc.inputIdx
		expected := tc.expectedOutput

		if actual := DropAt(arr, index); !slices.Equal(actual, expected) {
			t.Errorf("DropAt(%v,%d) = %v, expected %v", arr, index, actual, expected)
		}
	}
}

func TestInsertAt(t *testing.T) {
	testCases := []struct {
		inputArr       []string
		inputIdx       int
		InputElem      string
		expectedOutput []string
	}{
		{
			inputArr:       nil,
			inputIdx:       0,
			InputElem:      "hello",
			expectedOutput: []string{"hello"},
		},
		{
			inputArr:       []string{"hello"},
			inputIdx:       0,
			InputElem:      "goodbye",
			expectedOutput: []string{"goodbye", "hello"},
		},
		{
			inputArr:       []string{"hello"},
			inputIdx:       1,
			InputElem:      "goodbye",
			expectedOutput: []string{"hello", "goodbye"},
		},
		{
			inputArr:       []string{"apple", "orange", "pear", "pineapple"},
			inputIdx:       0,
			InputElem:      "apple",
			expectedOutput: []string{"apple", "apple", "orange", "pear", "pineapple"},
		},
		{
			inputArr:       []string{"apple", "orange", "pear", "pineapple"},
			inputIdx:       1,
			InputElem:      "apple",
			expectedOutput: []string{"apple", "apple", "orange", "pear", "pineapple"},
		},
		{
			inputArr:       []string{"apple", "orange", "pear", "pineapple"},
			inputIdx:       2,
			InputElem:      "apple",
			expectedOutput: []string{"apple", "orange", "apple", "pear", "pineapple"},
		},
		{
			inputArr:       []string{"apple", "orange", "pear", "pineapple"},
			inputIdx:       3,
			InputElem:      "apple",
			expectedOutput: []string{"apple", "orange", "pear", "apple", "pineapple"},
		},
		{
			inputArr:       []string{"apple", "orange", "pear", "pineapple"},
			inputIdx:       4,
			InputElem:      "apple",
			expectedOutput: []string{"apple", "orange", "pear", "pineapple", "apple"},
		},
	}

	for _, tc := range testCases {
		arr := tc.inputArr
		index := tc.inputIdx
		elem := tc.InputElem
		expected := tc.expectedOutput

		if actual := InsertAt(arr, index, elem); !slices.Equal(actual, expected) {
			t.Errorf("InsertAt(%v,%d,%s) = %v, expected %v", arr, index, elem, actual, expected)
		}
	}
}
