package floatvalidation

func ValidateFloat(value float64, min float64, max float64, precision int) (ok bool) {
	power_of_ten := 1
	for i := 0; i < precision; i++ {
		power_of_ten *= 10
	}
	corrected_value := float64(int(value * float64(power_of_ten)))
	final_float := corrected_value / float64(power_of_ten)

	ok = true
	if value > max || value < min || final_float != value {
		ok = false
	}

	return ok
}
