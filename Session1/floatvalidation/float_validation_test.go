package floatvalidation

import "testing"

/*
without using any library (including math), implement a function that validates a float64 value:
- having min and max, value must fall in range [min,max].
- having precision, value must not have more decimal points.
*/

type testCase struct {
	inputValue float64
	inputMin   float64
	inputMax   float64
	inputPrc   int
	output     bool
}

func TestValidateFloat(t *testing.T) {
	testCases := []testCase{
		{
			inputValue: 0.,
			inputMin:   0,
			inputMax:   1,
			inputPrc:   0,
			output:     true,
		},
		{
			inputValue: 1.,
			inputMin:   0,
			inputMax:   1,
			inputPrc:   0,
			output:     true,
		},
		{
			inputValue: -0.000001,
			inputMin:   0,
			inputMax:   1,
			inputPrc:   0,
			output:     false,
		},
		{
			inputValue: 1.000001,
			inputMin:   0,
			inputMax:   1,
			inputPrc:   0,
			output:     false,
		},
		{
			inputValue: 10.123,
			inputMin:   -11,
			inputMax:   11,
			inputPrc:   3,
			output:     true,
		},
		{
			inputValue: 10.1231,
			inputMin:   -11,
			inputMax:   11,
			inputPrc:   3,
			output:     false,
		},
		{
			inputValue: -15.000001,
			inputMin:   -20,
			inputMax:   -15,
			inputPrc:   5,
			output:     false,
		},
		{
			inputValue: -15.000001,
			inputMin:   -20,
			inputMax:   -15,
			inputPrc:   6,
			output:     true,
		},
	}

	for _, tc := range testCases {
		value := tc.inputValue
		min := tc.inputMin
		max := tc.inputMax
		precision := tc.inputPrc
		expected := tc.output

		if actual := ValidateFloat(value, min, max, precision); actual != expected {
			t.Errorf("ValidateFloat(%.6f,%.2f,%.2f,%d) = %t, expected %t", value, min, max, precision, actual, expected)
		}
	}
}
