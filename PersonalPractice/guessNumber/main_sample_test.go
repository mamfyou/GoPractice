package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SampleGame struct {
	number int
}

func MakeSampleGame(n int) *SampleGame {
	return &SampleGame{
		number: n,
	}
}

func (game *SampleGame) CheckNumber(g int) string {
	if g == game.number {
		return "CORRECT"
	} else if g < game.number {
		return "My Number is Greater"
	} else {
		return "My Number is Lower"
	}
}

func TestSampleGuessNumber(t *testing.T) {
	guess := 5
	game := MakeSampleGame(guess)
	res := GuessMyNumber(game)
	assert.Equal(t, fmt.Sprintf("Your Number was %d", guess), res, "You didn't Guessed it correctly")
}
