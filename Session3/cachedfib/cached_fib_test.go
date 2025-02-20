package cachedfib

import "testing"

/*
implement a cached recursive fibonacci function
*/

type testCase struct {
	input  int
	output int64
}

func TestCachedFib(t *testing.T) {
	testCases := []testCase{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
		{20, 6765},
		{30, 832040},
		{40, 102334155},
		{50, 12586269025},   // timeout without cache
		{60, 1548008755920}, // timeout without cache
	}
	
	fib := CachedFib()

	for _, tc := range testCases {
		n := tc.input
		expected := tc.output

		if actual := fib(n); actual != expected {
			t.Errorf("CachedFib()(%d) = %d, expected %d", n, actual, expected)
		}
	}
}
