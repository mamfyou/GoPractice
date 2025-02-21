package keyboard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample1(t *testing.T) {
	keyboard := NewKeyboard(1)
	output := keyboard.Enter("salam")
	assert.Equal(t, "salm", output)
}

func TestSample2(t *testing.T) {
	keyboard := NewKeyboard(2)
	output := keyboard.Enter("Welcome to CodeCup7")
	assert.Equal(t, "Welcome to Cdup7", output)
}
