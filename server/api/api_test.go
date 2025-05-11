package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/castlele/goaccounter/server/domain/accounts"
	"github.com/castlele/goaccounter/server/domain/user"
	"github.com/castlele/goaccounter/server/storage/database"
	"github.com/castlele/goaccounter/server/storage/models"
	"github.com/castlele/goaccounter/server/utils/testutils"
)

func TestRegistration(t *testing.T) {
	ts := *testutils.NewSuite(t)

	ts.Parallel("Valid user data", func(t *testing.T) {
		ts := *testutils.NewSuite(t)
		api := createAPI()
		newUser := models.UserData{
			Name:     "Javie",
			Password: "12345678",
		}

		r := createRequest(newUser, http.MethodPost, RouteRegister).sendRequest(api)

		ts.AssertStatusCode(http.StatusCreated, r.Code)
	})

	ts.Parallel("Invalid password len", func(t *testing.T) {
		ts := *testutils.NewSuite(t)
		api := createAPI()
		newUser := models.UserData{
			Name:     "Javie",
			Password: "1234",
		}

		r := createRequest(newUser, http.MethodPost, RouteRegister).sendRequest(api)

		ts.AssertStatusCode(http.StatusBadRequest, r.Code)
	})
}

func TestAccountCreation(t *testing.T) {
	ts := *testutils.NewSuite(t)

	ts.Parallel("Default balance", func(t *testing.T) {
		ts := *testutils.NewSuite(t)
		user := models.UserData{
			Name:     "Javie",
			Password: "12345678",
		}
		api := createAPI()
		jwt := register(api, user)
		newAccount := models.Account{
			Name: "Test Account",
		}

		r := createRequestWithAuth(newAccount, http.MethodPost, RouteAccounts, jwt).sendRequest(api)

		ts.AssertStatusCode(http.StatusCreated, r.Code)
	})

	ts.Parallel("Valid account data", func(t *testing.T) {
		ts := *testutils.NewSuite(t)
		user := models.UserData{
			Name:     "Javie",
			Password: "12345678",
		}
		api := createAPI()
		jwt := register(api, user)
		newAccount := models.Account{
			Name:    "Test Account",
			Balance: 1000,
		}

		r := createRequestWithAuth(newAccount, http.MethodPost, RouteAccounts, jwt).sendRequest(api)

		ts.AssertStatusCode(http.StatusCreated, r.Code)
	})

	ts.Parallel("Invalid account data", func(t *testing.T) {
		ts := *testutils.NewSuite(t)
		user := models.UserData{
			Name:     "Javie",
			Password: "12345678",
		}
		api := createAPI()
		jwt := register(api, user)
		newAccount := models.Account{
			Name:    "",
			Balance: -100,
		}

		r := createRequestWithAuth(newAccount, http.MethodPost, RouteAccounts, jwt).sendRequest(api)

		ts.AssertStatusCode(http.StatusBadRequest, r.Code)
	})
}

type Request struct {
	Recorder *httptest.ResponseRecorder
	Request  *http.Request
}

func register(api *API, user models.UserData) string {
	r := createRequest(user, http.MethodPost, RouteRegister).sendRequest(api)

	var registeredUser models.User

	if err := json.NewDecoder(r.Body).Decode(&registeredUser); err != nil {
		panic(err)
	}

	return registeredUser.JWT
}

func createRequest(message any, method string, route string) *Request {
	return createRequestWithAuth(message, method, route, "")
}

func createRequestWithAuth(message any, method string, route string, jwt string) *Request {
	body, _ := json.Marshal(message)

	r := httptest.NewRequest(method, route, bytes.NewBuffer(body))

	if jwt != "" {
		r.Header.Set("Authorization", "Bearer "+jwt)
	}

	return &Request{
		Request:  r,
		Recorder: httptest.NewRecorder(),
	}
}

func (r *Request) sendRequest(api *API) *httptest.ResponseRecorder {
	api.Router.ServeHTTP(r.Recorder, r.Request)

	return r.Recorder
}

func createAPI() *API {
	db := database.New()

	return New(
		user.New(db),
		accounts.New(db),
	)
}
