package roman

const (
	Invlaid   = -1
	None      = 0
	Equal     = 1
	Prefix    = 2
	Suffix    = 3
	Substring = 4
)

func ToRoman(n int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""
	for i := 0; i < len(values); i++ {
		count := int(n / values[i])
		for j := 0; j < count; j++ {
			result += symbols[i]
		}
		n %= values[i]
	}

	return result
}
