package primes

import "math"

func countPrime(num int) (count int) {
	for i := 2; i <= num; i++ {
		if isPrime(i) {
			count++
		}
	}

	return count
}

func isPrime(num int) bool {
	sqrt := math.Sqrt(float64(num))

	for i := 2; i <= int(sqrt); i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}