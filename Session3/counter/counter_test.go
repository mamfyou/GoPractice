package counter

import (
	"slices"
	"testing"
)

/*
implement a counter with custom step size using function closure
*/

const numCalls = 10

type testCase struct {
	stepSize int
	sequence []int
	hasError bool
	errorMsg string
}

func TestCounter(t *testing.T) {
	testCases := []testCase{
		{
			stepSize: -2,
			hasError: true,
			errorMsg: "step should be positive; got: -2",
		},
		{
			stepSize: -1,
			hasError: true,
			errorMsg: "step should be positive; got: -1",
		},
		{
			stepSize: 0,
			hasError: true,
			errorMsg: "step should be positive; got: 0",
		},
		{
			stepSize: 1,
			sequence: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			stepSize: 2,
			sequence: []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18},
		},
		{
			stepSize: 3,
			sequence: []int{0, 3, 6, 9, 12, 15, 18, 21, 24, 27},
		},
		{
			stepSize: 4,
			sequence: []int{0, 4, 8, 12, 16, 20, 24, 28, 32, 36},
		},
	}

	for _, tc := range testCases {
		stepSize := tc.stepSize
		expectedSeq := tc.sequence
		hasError := tc.hasError
		expectedErrMsg := tc.errorMsg

		counterFunc, err := Counter(stepSize)
		if err != nil {
			if hasError {
				if errMsg := err.Error(); errMsg == expectedErrMsg {
					continue
				} else {
					t.Errorf("error of Counter(%d) = %q, expected %q", stepSize, errMsg, expectedErrMsg)
				}
			} else {
				t.Errorf("error of Counter(%d) = %q, expected nil", stepSize, err)
			}
		} else {
			if hasError {
				t.Errorf("error of Counter(%d) = nil, expected %q", stepSize, expectedErrMsg)
			} else {
				actualSeq := generate(counterFunc)
				if !slices.Equal(actualSeq, expectedSeq) {
					t.Errorf("generated result of Counter(%d) = %v, expected %v", stepSize, actualSeq, expectedSeq)
				}
			}
		}
	}
}

/* ---------- helpers ---------- */

func generate(gen func() int) (result []int) {
	for range numCalls {
		value := gen()
		result = append(result, value)
	}

	return result
}
