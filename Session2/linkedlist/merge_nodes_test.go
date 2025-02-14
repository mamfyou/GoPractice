package linkedlist

import (
	"slices"
	"testing"
)

/*
implement a function that merges nodes in between zeros in a linked list:
	- function takes the head of linked list.
	- it merges consecutive nodes that have non-zero value into a single node (using sum).
	- the modified linked list should not contain any node with zero value.
*/

type testCase struct {
	input  *Node
	output *Node
}

func TestMergeNodes(t *testing.T) {
	testCases := []testCase{
		{
			input:  nil, // []
			output: nil, // []
		},
		{
			input:  &Node{Value: 0}, // [0]
			output: nil,             // []
		},
		{
			input:  &Node{Value: 0, Next: &Node{Value: 0}}, // [0 0]
			output: nil,                                    // []
		},
		// --------------------
		{
			input:  &Node{Value: 1}, // [1]
			output: &Node{Value: 1}, // [1]
		},
		{
			input:  &Node{Value: 0, Next: &Node{Value: 1}}, // [0 1]
			output: &Node{Value: 1},                        // [1]
		},
		{
			input:  &Node{Value: 1, Next: &Node{Value: 0}}, // [1 0]
			output: &Node{Value: 1},                        // [1]
		},
		// --------------------
		{
			input:  &Node{Value: 0, Next: &Node{Value: 1, Next: &Node{Value: 1}}}, // [0 1 1]
			output: &Node{Value: 2},                                               // [2]
		},
		{
			input:  &Node{Value: 1, Next: &Node{Value: 1, Next: &Node{Value: 0}}}, // [1 1 0]
			output: &Node{Value: 2},                                               // [2]
		},
		{
			input:  &Node{Value: 1, Next: &Node{Value: 0, Next: &Node{Value: 1}}}, // [1 0 1]
			output: &Node{Value: 1, Next: &Node{Value: 1}},                        // [1 1]
		},
		// --------------------
		{
			input:  &Node{Value: 0, Next: &Node{Value: 5, Next: &Node{Value: 10, Next: &Node{Value: 0}}}}, // [0 5 10 0]
			output: &Node{Value: 15},                                                                      // [15]
		},
		{
			input:  &Node{Value: 0, Next: &Node{Value: 5, Next: &Node{Value: 10, Next: &Node{Value: 0}}}}, // [0 5 10 0]
			output: &Node{Value: 15},                                                                      // [15]
		},
		{
			input:  &Node{Value: 0, Next: &Node{Value: 1, Next: &Node{Value: 0, Next: &Node{Value: 3, Next: &Node{Value: 0, Next: &Node{Value: 2, Next: &Node{Value: 2, Next: &Node{Value: 0}}}}}}}}, // [0 1 0 3 0 2 2 0]
			output: &Node{Value: 1, Next: &Node{Value: 3, Next: &Node{Value: 4}}},                                                                                                                    // [1 3 4]
		},
		{
			input:  &Node{Value: 0, Next: &Node{Value: 3, Next: &Node{Value: 1, Next: &Node{Value: 0, Next: &Node{Value: 4, Next: &Node{Value: 5, Next: &Node{Value: 2, Next: &Node{Value: 0}}}}}}}}, // [0 3 1 0 4 5 2 0]
			output: &Node{Value: 4, Next: &Node{Value: 11}},                                                                                                                                          // [4 11]
		},
	}

	for _, tc := range testCases {
		head := tc.input
		expected := tc.output
		actual := MergeNodes(head)

		inputs := toSlice(head)
		expecteds := toSlice(expected)
		actuals := toSlice(actual)

		if !slices.Equal(actuals, expecteds) {
			t.Errorf("MergeNodes(%v) = %v, expected %v", inputs, actuals, expecteds)
		}
	}
}

/* ---------- helpers ---------- */

func toSlice(head *Node) (res []int) {
	for cur := head; cur != nil; cur = cur.Next {
		res = append(res, cur.Value)
	}

	return res
}
