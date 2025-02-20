package cachedfib

var cache = make(map[int]int64)

func CachedFib() func(int) int64 {
	return func(number int) (result int64) {
		if number == 0 {
			return 0
		} else if number < 2 {
			return 1
		} else {
			if cache[number] != 0 {
				return cache[number]
			} else {
				cache[number] = CachedFib()(number-1) + CachedFib()(number-2)
			}
			return cache[number]
		}
	}
}
