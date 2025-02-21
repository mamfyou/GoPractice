package keyboard

const shift_chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ?!"

type Keyboard struct {
	Dure   int
	KeyMap map[string]int
}

func isShiftChar(word rune) bool {
	for _, v := range shift_chars {
		if v == word {
			return true
		}
	}

	return false
}

func canTypeChar(char rune, keyboard *Keyboard, is_in_shift bool) bool {
	is_shift_exceeded := is_in_shift && keyboard.KeyMap["shift"] >= keyboard.Dure
	is_char_exceeded := keyboard.KeyMap[string(char)] >= keyboard.Dure
	is_space_exceeded := string(char) == " " && keyboard.KeyMap["space"] >= keyboard.Dure
	return !(is_shift_exceeded || is_char_exceeded || is_space_exceeded)
}

func NewKeyboard(dure int) *Keyboard {
	return &Keyboard{
		Dure:   dure,
		KeyMap: make(map[string]int),
	}
}

func (keyboard *Keyboard) Enter(inp string) string {
	final_str := ""
	for _, v := range inp {
		is_shift_char := isShiftChar(v)
		if !canTypeChar(v, keyboard, is_shift_char) {
			continue
		}

		if is_shift_char {
			keyboard.KeyMap["shift"] += 1
		} else if string(v) == " " {
			keyboard.KeyMap["space"] += 1
		} else {
			keyboard.KeyMap[string(v)] += 1
		}

		final_str += string(v)
	}
	return final_str
}
