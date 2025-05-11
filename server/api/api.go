package api

import (
	"encoding/json"
	"net/http"

	"github.com/castlele/goaccounter/server/domain"
	"github.com/castlele/goaccounter/server/storage/models"
	"github.com/gorilla/mux"
)

type API struct {
	Router *mux.Router

	userInteractor    domain.UserInteractor
	accountInteractor domain.AccountInteractor
}

const (
	RouteRegister = "/register"
	RouteLogin    = "/login"

	RouteAccounts = "/accounts"
)

func New(userInteractor domain.UserInteractor, accountInteractor domain.AccountInteractor) *API {
	api := API{
		Router:            mux.NewRouter(),
		userInteractor:    userInteractor,
		accountInteractor: accountInteractor,
	}

	api.registerRoutes()

	return &api
}

func (api *API) registerRoutes() {
	api.Router.HandleFunc(RouteRegister, api.registrationHandler).Methods(http.MethodPost)
	api.Router.HandleFunc(RouteLogin, api.loginHandler).Methods(http.MethodPost)

	authRouter := api.Router.PathPrefix("/").Subrouter()
	authRouter.Use(AuthMiddleware)

	authRouter.HandleFunc(RouteAccounts, api.accountsHandler).Methods(http.MethodPost)
}

func (api *API) registrationHandler(w http.ResponseWriter, r *http.Request) {
	var userData models.UserData

	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		api.respondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		user, err := api.userInteractor.Register(userData)

		if err != nil {
			api.respondWithError(w, http.StatusBadRequest, err.Error())
		} else {
			api.respondWithJson(w, http.StatusCreated, user)
		}
	}
}

func (api *API) loginHandler(w http.ResponseWriter, r *http.Request) {
	var userData models.UserData

	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		api.respondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		user, err := api.userInteractor.Login(userData)

		if err != nil {
			api.respondWithError(w, http.StatusBadRequest, err.Error())
		} else {
			api.respondWithJson(w, http.StatusOK, user)
		}
	}
}

func (api *API) accountsHandler(w http.ResponseWriter, r *http.Request) {
	var account models.Account

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		api.respondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		err = api.accountInteractor.CreateAccount(account)

		if err != nil {
			api.respondWithError(w, http.StatusBadRequest, err.Error())
		} else {
			api.respondWithJson(w, http.StatusCreated, account)
		}
	}
}

func (api *API) respondWithError(w http.ResponseWriter, code int, message string) {
	api.respondWithJson(w, code, models.Error{Code: code, Message: message})
}

func (api *API) respondWithJson(w http.ResponseWriter, code int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(body)
}
