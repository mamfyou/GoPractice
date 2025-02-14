package pricecalculation

func CalculatePrice(prices map[string]int, cart []string) (price int) {
    total := 0
	for _, key := range cart {
		total += prices[key]
	}
	
	return total
}