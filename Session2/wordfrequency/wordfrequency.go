package wordfrequency

import "strings"

func WordFrequency(text string) map[string]int {
	words_map := make(map[string]int)

	words_list := strings.Fields(text)

	for _, word := range words_list {
		words_map[word] += 1
	}

	return words_map
}
