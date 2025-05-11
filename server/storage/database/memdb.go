package database

import (
	"github.com/castlele/goaccounter/server/domain"
	"github.com/castlele/goaccounter/server/storage/models"
)

type LocalDB struct {
	users []models.User
	accounts map[string][]models.Account
}

func New() *LocalDB {
	return &LocalDB{}
}

func (db *LocalDB) NewUser(name string, password string, jwt *domain.JWT) {
	db.users = append(db.users, models.User{
		Name:     name,
		Password: password,
		JWT:      jwt.Key,
		Exp:      jwt.Exp,
	})
}

func (db *LocalDB) UpdateUser(name string, user *models.User) {
	for i, user := range db.users {
		if user.Name == name {
			db.users[i] = models.User{
				Name:     name,
				Password: user.Password,
				JWT:      user.JWT,
				Exp:      user.Exp,
			}
			return
		}
	}
}

func (db *LocalDB) GetUser(name string) *models.User {
	for _, user := range db.users {
		if user.Name == name {
			return &user
		}
	}

	return nil
}

func (db *LocalDB) GetUsers() []models.User {
	return db.users
}

func (db *LocalDB) SaveAccount(account models.Account, userName string) {
	accounts := db.accounts[userName]

	accounts = append(accounts, account)
	// db.accounts[userName] = accounts
}

func (db *LocalDB) GetAccounts(userName string) []models.Account {
	return db.accounts[userName]
}
