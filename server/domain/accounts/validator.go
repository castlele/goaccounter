package accounts

import (
	"errors"

	"github.com/castlele/goaccounter/server/storage/models"
)

type validator struct{}

func NewValidator() *validator {
	return &validator{}
}

func (v *validator) ValidateAccount(account models.Account) error {
	var err error

	if err = v.ValidateAccountName(account.Name); err != nil {
		return err
	}

	if err = v.ValidateAccountBalance(account.Balance); err != nil {
		return err
	}

	return nil
}

func (v *validator) ValidateAccountName(name string) error {
	if len(name) == 0 {
		return errors.New("Account name can not be empty")
	}
	return nil
}

func (v *validator) ValidateAccountBalance(balance float64) error {
	if balance < 0 {
		return errors.New("Account balance cannot be negative")
	}
	return nil
}
