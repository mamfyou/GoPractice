package inversion

import (
	"slices"
	"testing"
)

/*
implement a function that inverts a map (swaps keys and values).
because multiple input keys can have the same value, store the original keys in a sorted slice.
*/

type testCase struct {
	input  map[string]string
	output map[string][]string
}

func TestInvertMap(t *testing.T) {
	testCases := []testCase{
		{
			input:  nil,
			output: map[string][]string{},
		},
		{
			input: map[string]string{
				"Alice":   "Engineering",
				"Bob":     "Marketing",
				"Charlie": "Engineering",
				"Diana":   "Sales",
			},
			output: map[string][]string{
				"Engineering": {"Alice", "Charlie"},
				"Marketing":   {"Bob"},
				"Sales":       {"Diana"},
			},
		},
		{
			input: map[string]string{
				"A": "X",
				"B": "Y",
				"C": "X",
			},
			output: map[string][]string{
				"X": {"A", "C"},
				"Y": {"B"},
			},
		},
		{
			input: map[string]string{
				"K": "X",
				"J": "X",
				"I": "X",
				"H": "X",
				"G": "X",
				"F": "X",
				"E": "X",
				"D": "X",
				"C": "X",
				"B": "X",
				"A": "X",
			},
			output: map[string][]string{
				"X": {"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K"},
			},
		},
		{
			input: map[string]string{
				"بوتکمپ 1": "GO",
				"بوتکمپ 2": "GO",
				"سلام 1":   "دنیا",
				"سلام 2":   "دنیا",
				"Hello 1":  "World!",
				"Hello 2":  "World!",
			},
			output: map[string][]string{
				"GO":     {"بوتکمپ 1", "بوتکمپ 2"},
				"دنیا":   {"سلام 1", "سلام 2"},
				"World!": {"Hello 1", "Hello 2"},
			},
		},
	}

	for _, tc := range testCases {
		input := tc.input
		expected := tc.output

		if actual := InvertMap(input); !mapsEqual(actual, expected) {
			t.Errorf("InvertMap(%v) = %v, expected %v", input, actual, expected)
		}
	}
}

/* ---------- helpers ---------- */

func mapsEqual(m1, m2 map[string][]string) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || !slices.Equal(v2, v1) {
			return false
		}
	}

	return true
}
