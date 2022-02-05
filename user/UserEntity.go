package user

import "backend/account"

type User struct {
	UserID   uint64
	Name     string `json:"name,omitempty"`
	Age      int8   `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Accounts []account.Account
}

type Users struct {
	Users []User
}
