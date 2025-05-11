package domain

import (
	"github.com/castlele/goaccounter/server/storage/models"
)

type UserInteractor interface {
	Register(user models.UserData) (*models.User, error)
	Login(user models.UserData) (*models.User, error)
}

type AccountInteractor interface {
	CreateAccount(account models.Account, userName string) error
	GetAccounts(userName string) []models.Account
}
