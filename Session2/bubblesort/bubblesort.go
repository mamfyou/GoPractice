package bubblesort

func Sort(arr []string) {
	if len(arr) < 2 {
		return
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

}
