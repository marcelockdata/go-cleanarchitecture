package domain

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid username or password")
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
