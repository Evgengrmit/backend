package user

import "backend/account"

type User struct {
	UserID         uint64
	Name           string            `json:"name"`
	Age            int8              `json:"age,omitempty"`
	Login          string            `json:"login"`
	Email          string            `json:"email"`
	HashedPassword []byte            `json:"-"`
	Accounts       []account.Account `json:"accounts,omitempty"`
}

type Users struct {
	Users []User
}
