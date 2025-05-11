package storage

import (
	"github.com/castlele/goaccounter/server/domain"
	"github.com/castlele/goaccounter/server/storage/models"
)

type Storage interface {
	NewUser(name string, password string, jwt *domain.JWT)
	GetUser(name string) *models.User
	UpdateUser(name string, user *models.User)
	GetUsers() []models.User

	SaveAccount(account models.Account, userName string)
	GetAccounts(userName string) []models.Account
}
