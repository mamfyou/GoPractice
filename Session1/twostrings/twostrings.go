package twostrings

const (
    Invlaid   = -1
    None      = 0
    Equal     = 1
    Prefix    = 2
    Suffix    = 3
    Substring = 4
)

func containsSubstring(source string, target string) (ok bool){
	rune_source := []rune(source)
	rune_target := []rune(target)
	
	len_rune_source := len(rune_source)
	len_rune_target := len(rune_target)

	for i := 0; i < len_rune_source; i++ {
		if i + len_rune_target > len_rune_source {
			return ok
		} else if string(rune_source[i:i + len_rune_target]) == target {
			ok = true
			break
		}
	}

	return ok
}

func isEqual(s1 string, s2 string) (ok bool) {
	rune_s1 := []rune(s1)
	rune_s2 := []rune(s2)

	if len(s1) == len(s2) {
		are_equal := true
		for i := 0; i < len(rune_s1); i++ {
			if !(rune_s1[i] == rune_s2[i]) {
				are_equal = false
				break
			}
		}
		if are_equal {
			return true
		}
	}

	return false
}

func Process(s1 string, s2 string) int {
	rune_s1 := []rune(s1)
	rune_s2 := []rune(s2)

	if (len(rune_s2) > len(rune_s1) || len(s1) == 0 || len(s2) == 0) {
		return Invlaid
	} 

    if isEqual(s1, s2) {
		return Equal
	}

	if string(rune_s1[:len(rune_s2)]) == s2 {
		return Prefix
	}

	if string(rune_s1[len(rune_s1) - len(rune_s2):]) == s2 {
		return Suffix
	}
	
	if containsSubstring(s1, s2) {
		return Substring
	}
	
	return 0

}