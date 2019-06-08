package users

import (
	"time"
)

type User struct {
	Username string `json:"username"`
	Group    string `json:"group"`
}

type Token struct {
	ID       int       `json:"id"`
	Group    string    `json:"group"`
	Key      string    `json:"key"`
	IsActive bool      `json:"is_active"`
	Created  time.Time `json:"created"`
	Expires  time.Time `json:"expires"`
}

type ChangePasswordResponse struct {
	Success bool `json:"success"`
}

type CreateTokenResponse struct {
	User  User  `json:"user"`
	Token Token `json:"token"`
}
