package dataset

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

/*
normalize columns of a csv dataset:
	- normalization: scaling data into interval [0,1]
	- norm(D_ij) = (D_ij - min_j) / (max_j - min_j)
	- example: col1 = [1,2,5]; norm(col1) = [(1-1)/(5-1),(2-1)/(5-1),(5-1)/(5-1)]=[0,0.25,1]

- read/parse csv file
- normalize its data
- OVERWRITE the result back to the csv file
*/

const fileAddress = "data.csv"

type testCase struct {
	input    []string
	output   []string
	hasError bool
}

func TestPreprocessDataset(t *testing.T) {
	testCases := []testCase{
		{
			input: []string{
				"1,2",
				"3,4",
				"5,hello", // "hello" can't be converted to float64
			},
			hasError: true,
		},
		{
			input: []string{
				"1",
				"1,2", // inconsistent column count in each row
				"3",
			},
			hasError: true,
		},
		// ----------
		{
			input:  []string{},
			output: []string{},
		},
		{
			input: []string{
				"-1,8",
			},
			output: []string{
				"NaN,NaN",
			},
		},
		{
			input: []string{
				"1",
				"2",
			},
			output: []string{
				"0.00",
				"1.00",
			},
		},
		{
			input: []string{
				"1,4",
				"2,5",
				"3,6",
			},
			output: []string{
				"0.00,0.00",
				"0.50,0.50",
				"1.00,1.00",
			},
		},
		{
			input: []string{
				"5,4,3",
				"3,4,5",
				"1,3,2",
				"0,8,9",
			},
			output: []string{
				"1.00,0.20,0.14",
				"0.60,0.20,0.43",
				"0.20,0.00,0.00",
				"0.00,1.00,1.00",
			},
		},
	}

	testFunc := func(tc testCase) (err error) {
		inputLines := tc.input
		expectedLines := tc.output
		hasError := tc.hasError

		err = createTestFile(inputLines)
		if err != nil {
			return err
		}

		defer func() {
			_ = os.Remove(fileAddress)
		}()

		err = PreProcessDataset(fileAddress)
		if err != nil {
			if hasError {
				return nil
			} else {
				return fmt.Errorf("expected nil error; got %q", err.Error())
			}
		} else if hasError {
			return fmt.Errorf("expected error; got nil")
		}

		return verifyTestFileContent(expectedLines)
	}

	for _, tc := range testCases {
		if err := testFunc(tc); err != nil {
			t.Fatal(err)
		}
	}
}

/* ---------- helpers ---------- */

func createTestFile(lines []string) (err error) {
	file, err := os.Create(fileAddress)
	if err != nil {
		return err
	}

	defer file.Close()

	for _, line := range lines {
		_, err = fmt.Fprintln(file, line)
		if err != nil {
			return err
		}
	}

	return nil
}

func verifyTestFileContent(lines []string) (err error) {
	file, err := os.Open(fileAddress)
	if err != nil {
		return err
	}

	defer file.Close()

	i := 0
	for fscan := bufio.NewScanner(file); fscan.Scan(); {
		expected := lines[i]
		actual := fscan.Text()
		i++

		if actual != expected {
			return fmt.Errorf("unexpected content at line %d: %q != %q", i, actual, expected)
		}
	}

	return nil
}
