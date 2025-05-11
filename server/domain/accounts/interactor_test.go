package accounts

import (
	"testing"
	"time"

	"github.com/castlele/goaccounter/server/domain"
	"github.com/castlele/goaccounter/server/storage/database"
	"github.com/castlele/goaccounter/server/storage/models"
	"github.com/castlele/goaccounter/server/utils/testutils"
)

func TestAccountCreation(t *testing.T) {
	ts := *testutils.NewSuite(t)

	ts.Parallel("Create new account", func(t *testing.T) {
		ts := *testutils.NewSuite(t)
		users := make([]models.User, 1)
		users = append(users, models.User{ Name: "Javie", Password: "Helloworld" })
		sut := createInteractor(users)

		err := sut.CreateAccount(models.Account{ Name: "Javie's Acc"}, "Javie")

		ts.AssertNoError(err)
		ts.AssertEquals(len(sut.GetAccounts("Javie")) == 1, "Wrong amount of accounts")
		ts.AssertEquals(sut.GetAccounts("Javie")[0].Name == "Javie's Acc", "Wrong name of the account")
	})
}

func createInteractor(users []models.User) domain.AccountInteractor {
	db := database.New()

	for _, user := range users {
		db.NewUser(user.Name, user.Password, &domain.JWT{ Key: "", Exp: time.Now() })
	}

	return New(db)
}
