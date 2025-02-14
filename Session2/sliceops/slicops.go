package sliceops

func RemoveAll(arr []string, key string) (res []string) {
	new_array := make([]string, 0)
	for i := 0; i < len(arr); i++ {
		if arr[i] != key {
			new_array = append(new_array, arr[i])
		}
	}
	return new_array
}

func DropAt(arr []string, index int) []string {
	new_array := make([]string, len(arr)-1)
	for i := 0; i < index; i++ {
		new_array[i] = arr[i]
	}
	for j := index; j < len(new_array); j++ {
		new_array[j] = arr[j+1]
	}

	return new_array
}

func InsertAt(arr []string, index int, elem string) []string {
	new_array := make([]string, len(arr)+1)
	for i := 0; i < index; i++ {
		new_array[i] = arr[i]
	}
	new_array[index] = elem
	for j := index + 1; j < len(new_array); j++ {
		new_array[j] = arr[j-1]
	}

	return new_array
}
