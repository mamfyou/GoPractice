package pricecalculation

import "testing"

/*
implement a function that calculates total price of items in a 'cart'.
if price of an item is undefined, then it's free!
*/

type testCase struct {
	inputPrices map[string]int
	inputCart   []string
	output      int
}

func TestCalculatePrice(t *testing.T) {
	testCases := []testCase{
		{
			inputPrices: nil,
			inputCart:   nil,
			output:      0,
		},
		{
			inputPrices: map[string]int{},
			inputCart:   []string{},
			output:      0,
		},
		{
			inputPrices: map[string]int{
				"banana": 1,
				"apple":  2,
				"milk":   3,
			},
			inputCart: []string{"apple", "banana", "banana", "milk"},
			output:    7, // 2 + 1 + 1 + 3
		},
		{
			inputPrices: map[string]int{
				"banana": 1,
				"apple":  2,
				"milk":   3,
			},
			inputCart: []string{"apple", "chocolate", "milk"},
			output:    5, // 2 + 0 + 3
		},
		{
			inputPrices: map[string]int{
				"apple": 5,
				"milk":  4,
			},
			inputCart: []string{"apple", "apple", "apple", "banana", "banana", "chocolate", "milk", "milk"},
			output:    23, // 3*5 + 2*0 + 0 + 2*4
		},
	}

	for _, tc := range testCases {
		prices := tc.inputPrices
		cart := tc.inputCart
		expected := tc.output

		if actual := CalculatePrice(prices, cart); actual != expected {
			t.Errorf("CalculatePrice(%v,%v) = %d, expected %d", prices, cart, actual, expected)
		}
	}
}
