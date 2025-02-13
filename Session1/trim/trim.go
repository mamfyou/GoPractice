package trim

func TrimSpace(s string) string {
	rune_s := []rune(s)

	space_count := 0
	for _, c := range s {
		if string(c) == " " {
			space_count++
		}
	}

	if space_count == len(rune_s) {
		return ""
	}

	start_index := 0
	for _, c := range s {
		if string(c) == " " {
			start_index++
		} else {
			break
		}
	}

	end_index := len(rune_s) - 1
	for i := len(rune_s) - 1; i > 0; i-- {
		if string(rune_s[i]) == " " {
			end_index--
		} else {
			break
		}
	}

	return string(rune_s[start_index : end_index+1])
}
