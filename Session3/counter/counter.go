package counter

import (
	"fmt"
)

func Counter(step_size int) (fun func() int, err error) {
	counter := -step_size
	if step_size <= 0 {
		return nil, fmt.Errorf("step should be positive; got: %d", step_size)
	}
	return func() int {
		counter += step_size
		return counter
	}, err
}