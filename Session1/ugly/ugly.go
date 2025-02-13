package ugly

import "math"

func isPrime(num int) bool {
	sqrt := math.Sqrt(float64(num))

	for i := 2; i <= int(sqrt); i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func IsUgly(n int) (ok bool) {
	if n == 0 || n == 1 {
		return false
	}
	
	for i := 7; i <= n; i++ {
		if (n % i == 0 && isPrime(i)) {
			ok = false
			return ok
		}
	}
	ok = true
	
    return ok
}