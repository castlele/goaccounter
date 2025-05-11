package user

import (
	"testing"

	"github.com/castlele/goaccounter/server/storage/database"
	"github.com/castlele/goaccounter/server/storage/models"
	"github.com/castlele/goaccounter/server/utils/testutils"
)

func TestRegister(t *testing.T) {
	ts := testutils.NewSuite(t)

	ts.Parallel("Register user with valid fields", func(t *testing.T) {
		ts := testutils.NewSuite(t)
		user := models.UserData{
			Name:     "Javie",
			Password: "12345678",
		}
		sut := createInteractor()

		_, err := sut.Register(user)

		ts.AssertNoError(err)
		ts.AssertEquals(
			len(sut.GetUsers()) == 1,
			"Invalid users amount",
		)
	})

	ts.Parallel("Register user with short password field", func(t *testing.T) {
		ts := testutils.NewSuite(t)
		user := models.UserData{
			Name:     "Javie",
			Password: "1234567",
		}
		sut := createInteractor()

		_, err := sut.Register(user)

		ts.AssertError(err)
		ts.AssertEquals(
			len(sut.GetUsers()) == 0,
			"Invalid users amount",
		)
	})

	ts.Parallel("Register user that was previously registered", func(t *testing.T) {
		ts := testutils.NewSuite(t)
		user := models.UserData{
			Name:     "Javie",
			Password: "12345678",
		}
		sut := createInteractor()
		sut.Register(user)

		_, err := sut.Register(user)

		ts.AssertError(err)
		ts.AssertEquals(
			len(sut.GetUsers()) == 1,
			"Invalid users amount",
		)
	})
}

func TestLogin(t *testing.T) {
	ts := testutils.NewSuite(t)

	ts.Parallel("Login with unregistered user", func(t *testing.T) {
		ts := testutils.NewSuite(t)
		user := models.UserData{
			Name:     "Javie",
			Password: "12345678",
		}
		sut := createInteractor()

		result, err := sut.Login(user)

		ts.AssertError(err)
		ts.AssertEquals(
			err.Error() == ErrorNoUserFound,
			"Wrong error found",
		)
		ts.AssertNil(result)
		ts.AssertEquals(
			len(sut.GetUsers()) == 0,
			"Invalid users amount",
		)
	})

	ts.Parallel("Login with registered user", func(t *testing.T) {
		ts := testutils.NewSuite(t)
		user := models.UserData{
			Name:     "Javie",
			Password: "12345678",
		}
		sut := createInteractor()
		sut.Register(user)

		result, err := sut.Login(user)

		ts.AssertNoError(err)
		ts.AssertEquals(
			len(sut.GetUsers()) == 1,
			"Invalid users amount",
		)
		ts.AssertEquals(
			result.Name == user.Name && result.Password == user.Password,
			"Invalid users amount",
		)
	})

	ts.Parallel("Login with registered user and wrong password", func(t *testing.T) {
		ts := testutils.NewSuite(t)
		user := models.UserData{
			Name:     "Javie",
			Password: "12345678",
		}
		userWithWrongPas := user
		userWithWrongPas.Password = user.Password + "9"
		sut := createInteractor()
		sut.Register(user)

		result, err := sut.Login(userWithWrongPas)

		ts.AssertError(err)
		ts.AssertEquals(
			err.Error() == ErrorInvalidPassword,
			"Wrong error passed",
		)
		ts.AssertEquals(
			len(sut.GetUsers()) == 1,
			"Invalid users amount",
		)
		ts.AssertNil(result)
	})
}

func createInteractor() *UserInteractor {
	return New(database.New())
}
