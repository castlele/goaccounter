package user

import "errors"

type validator struct {
}

const (
	ErrorShortPassword = "Password is less than 8 symbols"
)

func NewValidator() *validator {
	return &validator{}
}

func (v *validator) ValidatePassword(pass string) error {
	if len(pass) < 8 {
		return errors.New(ErrorShortPassword)
	}

	return nil
}
