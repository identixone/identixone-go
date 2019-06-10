package users

import (
	"fmt"
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

type ChangePasswordRequest struct {
	Password1   string `json:"password1"`
	Password2   string `json:"password2"`
	ResetTokens bool   `json:"reset_tokens"`
}

func (cp ChangePasswordRequest) Validate() error {
	if cp.Password1 == "" || cp.Password2 == "" {
		return fmt.Errorf("Password1 and Password2 is required")
	}

	if cp.Password1 != cp.Password2 {
		return fmt.Errorf("passwords is different")
	}
	return nil
}

type ChangePasswordResponse struct {
	Success bool `json:"success"`
}

type CreateTokenRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Permanent bool   `json:"-"`
}

type CreateTokenResponse struct {
	User  User  `json:"user"`
	Token Token `json:"token"`
}
