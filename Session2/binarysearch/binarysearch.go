package binarysearch

func BinarySearch(arr []string, key string) int {
	return binarySearchHelper(arr, key, 0, len(arr)-1)
}

func binarySearchHelper(arr []string, key string, low, high int) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2

	if arr[mid] == key {
		return mid
	} else if arr[mid] < key {
		return binarySearchHelper(arr, key, mid+1, high)
	} else {
		return binarySearchHelper(arr, key, low, mid-1)
	}
}
