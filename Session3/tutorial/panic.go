package tutorial

import (
	"strconv"
	"strings"
)

func GetRowData(line string) (row []float64, err error) {
	values := strings.Split(line, ",")

	row = make([]float64, len(values))
	for i,v := range values {
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, err
		}
		row[i] = f
	}

	return row, nil
}