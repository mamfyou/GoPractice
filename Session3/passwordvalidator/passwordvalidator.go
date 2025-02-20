package passwordvalidator

import (
	"errors"
)

type PasswordMinLengthValidator struct {
	password string
	length   int
}

type PasswordHasNumberValidator struct {
	password string
}

type PasswordSpecialCharValidator struct {
	password string
}

var ErrPasswordIsTooShort = errors.New("password is too short")
var ErrPasswordContainsNoNumbers = errors.New("password contains no numbers")
var ErrPasswordContainsNoSpecialCharacters = errors.New("password contains no special characters")

func (p PasswordMinLengthValidator) Validate() error {
	if len([]rune(p.password)) < p.length {
		return ErrPasswordIsTooShort
	}
	return nil
}

func (p PasswordHasNumberValidator) Validate() error {
	for _, v := range p.password {
		for _, v2 := range "1234567890" {
			if v == v2 {
				return nil
			}
		}
	}
	return ErrPasswordContainsNoNumbers
}

func (p PasswordSpecialCharValidator) Validate() error {
	specialChars := "!@#$%^&*()"
	for _, v := range p.password {
		for _, v2 := range specialChars {
			if v == v2 {
				return nil
			}
		}
	}
	return ErrPasswordContainsNoSpecialCharacters
}

func NewPasswordMinLengthValidator(password string, length int) *PasswordMinLengthValidator {
	return &PasswordMinLengthValidator{password: password, length: length}
}

func NewPasswordNumberValidator(password string) *PasswordHasNumberValidator {
	return &PasswordHasNumberValidator{password: password}
}

func NewPasswordSpecialCharValidator(password string) *PasswordSpecialCharValidator {
	return &PasswordSpecialCharValidator{password: password}
}
