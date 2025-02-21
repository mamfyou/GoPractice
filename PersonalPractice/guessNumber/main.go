package main

import "fmt"

func Guess(game Game, low int, high int) (answer int) {
	mid_number := int((low + high) / 2)
	result := game.CheckNumber(mid_number)

	if result == "CORRECT" {
		return mid_number
	} else if result == "My Number is Greater" {
		return Guess(game, mid_number, high)
	} else if result == "My Number is Lower" {
		return Guess(game, low, mid_number)
	} else {
		return -1
	}
}

// 1-360 -> 180
// 1-180 -> 90
// 1-90 -> 45
// 1-45 -> 22
// 1-22 -> 11
// 1-11 -> 6
// 1-6 -> 3
// 3 - 6 -> 4
// 4 - 6 -> 5

func GuessMyNumber(game Game) string {
	_ = game.CheckNumber(0)
	return fmt.Sprintf("Your Number was %d", Guess(game, 1, 360))
}
