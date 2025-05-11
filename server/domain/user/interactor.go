package user

import (
	"errors"
	"time"

	"github.com/castlele/goaccounter/server/domain"
	"github.com/castlele/goaccounter/server/storage"
	"github.com/castlele/goaccounter/server/storage/models"
)

const (
	ErrorUserExists      = "User with this name already exists"
	ErrorNoUserFound     = "No user was found with the given name"
	ErrorInvalidPassword = "Invalid password was passed"
)

type UserInteractor struct {
	db        storage.Storage
	validator *validator
}

func New(db storage.Storage) *UserInteractor {
	return &UserInteractor{
		db:        db,
		validator: NewValidator(),
	}
}

func (i *UserInteractor) GetUsers() []models.User {
	return i.db.GetUsers()
}

func (i *UserInteractor) Register(user models.UserData) (*models.User, error) {
	err := i.validator.ValidatePassword(user.Password)

	if err != nil {
		return nil, err
	}

	if i.db.GetUser(user.Name) == nil {
		jwt, err := domain.GenerateJWT(user.Name, time.Now().Add(24*time.Hour))

		if err != nil {
			return nil, err
		}

		i.db.NewUser(
			user.Name,
			user.Password,
			jwt,
		)
	} else {
		return nil, errors.New(ErrorUserExists)
	}

	return i.db.GetUser(user.Name), nil
}

func (i *UserInteractor) Login(user models.UserData) (*models.User, error) {
	registeredUser := i.db.GetUser(user.Name)

	if registeredUser == nil {
		return nil, errors.New(ErrorNoUserFound)
	}

	if registeredUser.Password != user.Password {
		return nil, errors.New(ErrorInvalidPassword)
	}

	jwt, err := domain.GenerateJWT(user.Name, time.Now().Add(24*time.Hour))

	if err != nil {
		return nil, err
	}

	registeredUser.JWT = jwt.Key
	registeredUser.Exp = jwt.Exp

	i.db.UpdateUser(
		user.Name,
		registeredUser,
	)

	return registeredUser, nil
}
