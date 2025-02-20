package tutorial

type Predicate func(int) bool

func Filter(input []int, pred Predicate) (output []int) {
	for _, v := range input {
		if pred(v) {
			output = append(output, v)
		}
	}
	return output
}
