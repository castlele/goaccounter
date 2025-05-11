package accounts

import (
	"github.com/castlele/goaccounter/server/storage"
	"github.com/castlele/goaccounter/server/storage/models"
)

type accountInteractor struct {
	db        storage.Storage
	validator *validator
}

func New(db storage.Storage) *accountInteractor {
	return &accountInteractor{
		db:        db,
		validator: NewValidator(),
	}
}

func (i *accountInteractor) CreateAccount(account models.Account, userName string) error {
	if err := i.validator.ValidateAccountName(account.Name); err != nil {
		return err
	}

	i.db.SaveAccount(account, userName)

	return nil
}

func (i *accountInteractor) GetAccounts(userName string) []models.Account {
	return i.db.GetAccounts(userName)
}
