package wordfrequency

import "testing"

/*
implement a function that calculates frequency of each word in a text.
you might need to split the input text into words by space characters.
*/

type testCase struct {
	input  string
	output map[string]int
}

func TestWordFrequency(t *testing.T) {
	testCases := []testCase{
		{
			input:  "",
			output: map[string]int{},
		},
		{
			input:  "\t   \n",
			output: map[string]int{},
		},
		{
			input: "go is fun and go is easy",
			output: map[string]int{
				"go":   2,
				"is":   2,
				"fun":  1,
				"and":  1,
				"easy": 1,
			},
		},
		{
			input: "در این متن واژه متن دو بار و واژه واژه سه بار تکرار شده است",
			output: map[string]int{
				"در":    1,
				"این":   1,
				"متن":   2,
				"واژه":  3,
				"دو":    1,
				"بار":   2,
				"و":     1,
				"سه":    1,
				"تکرار": 1,
				"شده":   1,
				"است":   1,
			},
		},
	}

	for _, tc := range testCases {
		text := tc.input
		expected := tc.output

		if actual := WordFrequency(text); !mapsEqual(actual, expected) {
			t.Errorf("WordFrequency(%q) = %v, expected %v", text, actual, expected)
		}
	}
}

/* ---------- helpers ---------- */

func mapsEqual(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || v2 != v1 {
			return false
		}
	}

	return true
}
