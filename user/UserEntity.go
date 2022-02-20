package user

import (
	"backend/account"
)

type User struct {
	ID             int               `json:"id"`
	Name           string            `json:"name"`
	Age            int8              `json:"age,omitempty"`
	Login          string            `json:"login"`
	Email          string            `json:"email"`
	HashedPassword []byte            `json:"-"`
	Accounts       []account.Account `json:"accounts,omitempty"`
}
