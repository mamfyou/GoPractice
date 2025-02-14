package inversion

func IsInSlice(slice []string, key string) (ok bool) {
	for _, value := range slice {
		if value == key {
			return true
		}
	}

	return ok
}

func BubbleSortSlice(arr []string) (sorted []string) {
	if len(arr) < 2 {
		return arr
	}

	len_sorted_arr := len(arr)

	for i := 0; i < len_sorted_arr-1; i++ {
		flag := false
		for j := 0; j < len_sorted_arr-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}

	return arr
}

func InvertMap(input map[string]string) (output map[string][]string) {
	output = make(map[string][]string, 0)

	for key, value := range input {
		if !IsInSlice(output[value], key) {
			output[value] = append(output[value], key)
		}

	}

	for key, value := range output {
		output[key] = BubbleSortSlice(value)
	}

	return output
}
