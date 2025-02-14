package dynamicmatrix

import "math"

func SetElement(matrix [][]int, i int, j int, value int) (result [][]int) {
	row_count := int(math.Max(float64(i+1), float64(len(matrix))))
	horizontal := make([][]int, row_count)

	for m := 0; m < row_count; m++ {
		column_count := 0
		if len(matrix) > 0 {
			column_count = int(math.Max(float64(j+1), float64(len(matrix[0]))))
		} else {
			column_count = 1
		}
		horizontal[m] = make([]int, column_count)
	}

	for k := 0; k < len(matrix); k++ {
		copy(horizontal[k], matrix[k])
	}

	horizontal[i][j] = value

	return horizontal
}
