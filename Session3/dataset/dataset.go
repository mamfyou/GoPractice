package dataset

import (
	"os"
	"strconv"
	"strings"
	"bufio"
	"fmt"
)

func findMax(input []float64) (result float64) {
	max_value := 0.0
	for _, v := range input {
		if v > float64(max_value) {
			max_value = v
		}
	}
	return float64((max_value))
}
func findMin(input []float64) (result float64) {
	min_value := input[0]
	for _, v := range input {
		if v < min_value {
			min_value = v
		}
	}
	return float64((min_value))
}
func PreProcessDataset(address string) (err error) {
	file, err := os.Open(address)
	if err != nil {
		return err
	}

	defer file.Close()

	columns := make([][]float64, 0)

	i := 0
	previous_len := 0
	for fscan := bufio.NewScanner(file); fscan.Scan(); {
		actual := fscan.Text()
		if len(columns) == 0 {
			columnCount := len(strings.Split(actual, ","))
			for i := 0; i < columnCount; i++ {
				columns = append(columns, []float64{})
			}
		} else if previous_len != len(strings.Split(actual, ",")) {
			return fmt.Errorf("inconsistent column count in each row")
		}
		values := strings.Split(actual, ",")
		for j := 0; j < len(values); j++ {
			value, err := strconv.ParseFloat(values[j], 64)
			if err != nil {
				return fmt.Errorf("non-numeric value found: %v", values[j])
			}
			columns[j] = append(columns[j], value)
		}
		previous_len = len(values)
		i++
	}

	final_file := []byte{}
	for fscan := bufio.NewScanner(file); fscan.Scan(); {
		actual := fscan.Text()
		new_string := ""
		values := strings.Split(actual, ",")
		for j := 0; j < len(values); j++ {
			max_value := findMax(columns[j])
			min_value := findMin(columns[j])
			value, _ := strconv.ParseFloat(values[j], 64)
			final_value := (value - min_value) / (max_value - min_value)
			new_string += strconv.FormatFloat(final_value, 'f', -1, 64)

			if j != len(values)-1 {
				new_string += ","
			}
		}
		final_file = append(final_file, []byte(new_string+"\n")...)
		i++
	}
	err = os.WriteFile(address, final_file, 0644)
	if err != nil {
		return err
	}


	return nil
}

