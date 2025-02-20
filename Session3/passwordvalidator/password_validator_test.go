package passwordvalidator

import (
	"errors"
	"testing"
)

/*
implement three password validator types:
- PasswordMinLengthValidator: password must have a minimum number of characters.
- PasswordNumberValidator: password must have at least one number character.
- PasswordSpecialCharValidator: password must have at least one special character from `!@#$%^&*()`.
implementations must SATISFY the PasswordValidator interface.
*/

type PasswordValidator interface {
	Validate() error
}

func TestPasswordMinLengthValidator(t *testing.T) {
	testCases := []struct {
		password string
		length   int
		hasErr   bool
	}{
		{"", 0, false},
		{"", 1, true},
		{"🤣", 1, false},
		{"🤣", 2, true},
		{"🤣🤣", 2, false},
		{"🤣🤣", 3, true},
	}

	for _, tc := range testCases {
		password := tc.password
		length := tc.length
		hasErr := tc.hasErr

		var validator PasswordValidator = NewPasswordMinLengthValidator(password, length)

		err := validator.Validate()
		if err != nil {
			if !hasErr {
				t.Errorf("expected nil error for input=%q, length=%d; got %q", password, length, err.Error())
			} else if !errors.Is(err, ErrPasswordIsTooShort) {
				t.Errorf("expected error %q; got %q", ErrPasswordIsTooShort.Error(), err.Error())
			}
		} else if hasErr {
			t.Errorf("expected error %q; got nil", ErrPasswordIsTooShort.Error())
		}
	}
}

func TestPasswordNumberValidator(t *testing.T) {
	testCases := []struct {
		password string
		hasErr   bool
	}{
		{"password", true},
		{"password1", false},
		{"123456", false},
		{"pass1word", false},
		{"!@#$%^&*", true},
		{"🤣🤣1🤣", false},
		{"🤣🤣🤣", true},
	}

	for _, tc := range testCases {
		password := tc.password
		hasErr := tc.hasErr

		var validator PasswordValidator = NewPasswordNumberValidator(password)

		err := validator.Validate()
		if err != nil {
			if !hasErr {
				t.Errorf("expected nil error for input=%q; got %q", password, err.Error())
			} else if !errors.Is(err, ErrPasswordContainsNoNumbers) {
				t.Errorf("expected error %q; got %q", ErrPasswordContainsNoNumbers.Error(), err.Error())
			}
		} else if hasErr {
			t.Errorf("expected error %q; got nil", ErrPasswordContainsNoNumbers.Error())
		}
	}
}

func TestPasswordSpecialCharValidator(t *testing.T) {
	testCases := []struct {
		password string
		hasErr   bool
	}{
		{"password", true},
		{"password!", false},
		{"!@#$%^&*", false},
		{"pass@word", false},
		{"123456", true},
		{"🤣🤣🤣", true},
		{"🤣🤣🤣!", false},
		{"pass word", true},
		{"pass!word", false},
		{"pass@word", false},
		{"pass#word", false},
		{"pass$word", false},
		{"pass%%word", false},
		{"pass^word", false},
		{"pass&word", false},
		{"pass*word", false},
		{"pass(word", false},
		{"pass)word", false},
		{"pass.word", true},
	}

	for _, tc := range testCases {
		password := tc.password
		hasErr := tc.hasErr

		var validator PasswordValidator = NewPasswordSpecialCharValidator(password)

		err := validator.Validate()
		if err != nil {
			if !hasErr {
				t.Errorf("expected nil error for input=%q; got %q", password, err.Error())
			} else if !errors.Is(err, ErrPasswordContainsNoSpecialCharacters) {
				t.Errorf("expected error %q; got %q", ErrPasswordContainsNoSpecialCharacters.Error(), err.Error())
			}
		} else if hasErr {
			t.Errorf("expected error %q; got nil", ErrPasswordContainsNoSpecialCharacters.Error())
		}
	}
}
