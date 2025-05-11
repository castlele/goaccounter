package models

import "time"

// Passed by client to register new User
type UserData struct {
	Name     string `json:"user"`
	Password string `json:"password"`
}

type User struct {
	Name     string    `json:"user"`
	Password string    `json:"password"`
	JWT      string    `json:"jwt"`
	Exp      time.Time `json:"exp"`
}

type Account struct {
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
