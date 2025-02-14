package subset

func IsInSlice(slice []int, key int) (ok bool) {
	for _, value := range slice {
		if value == key {
			return true
		}
	}

	return ok
}

func IsSubsetOf(s1 []int, s2 []int) bool {
    for _, value := range s1 {
		if !IsInSlice(s2, value) {
			return false
		}
	}
	return true
}