package user

import (
	"backend/account"
	"backend/account/balance"
	"errors"
	"fmt"
)

func (u *User) FindAccount(id uint64) (*account.Account, error) {
	for i := range u.Accounts {
		if id != u.Accounts[i].AccountID {
			return &u.Accounts[i], nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Account %d not found", id))
}
func NewUser() *User {
	return &User{UserID: uint64(users.NumberOfUsers() + 1)}
}

func (u *User) GetUser() User {
	return *u

}
func (u *User) GetAccounts() []account.Account {
	return u.Accounts
}
func (u *User) TopUpAccount(id uint64, b balance.Balance) error {
	foundAcc, err := u.FindAccount(id)
	if err != nil {
		return err
	}
	err = foundAcc.TopUp(b)
	if err != nil {
		return err
	}
	return nil

}
func (u *User) TakeOffAccount(id uint64, b balance.Balance) error {
	foundAcc, err := u.FindAccount(id)
	if err != nil {
		return err
	}
	err = foundAcc.TakeOff(b)
	if err != nil {
		return err
	}
	return nil

}
func (u *User) CreateAccount(currency balance.Currency) {
	newAccount := account.Account{
		AccountID: uint64(len(accounts) + 1),
		UserID:    u.UserID,
		Balance:   balance.Balance{Currency: currency}}
	u.Accounts = append(u.Accounts, newAccount)
	accounts = append(accounts, newAccount)
}

// USERS

func (users *Users) AddUser(u User) {
	users.Users = append(users.Users, u)
}

func (users *Users) NumberOfUsers() int {
	return len(users.Users)
}

func (users *Users) FindUserByName(name string) (*User, error) {
	for i := range users.Users {
		if users.Users[i].Name == name {
			return &users.Users[i], nil
		}
	}
	return nil, errors.New(fmt.Sprintf("User %s not found", name))
}

func (users *Users) GetUsers() []User {
	return users.Users
}
func (users *Users) TopUpForUser(username string, accData *account.Account) (User, error) {
	foundUser, err := users.FindUserByName(username)
	if err != nil {
		return User{}, err
	}
	err = foundUser.TopUpAccount(accData.AccountID, accData.Balance)
	if err != nil {
		return User{}, err
	}
	return *foundUser, nil
}
func (users *Users) TakeOffForUser(username string, accData *account.Account) (User, error) {
	foundUser, err := users.FindUserByName(username)
	if err != nil {
		return User{}, err
	}
	err = foundUser.TakeOffAccount(accData.AccountID, accData.Balance)
	if err != nil {
		return User{}, err
	}
	return *foundUser, nil
}
func (users *Users) TransferBetweenUsers(senderName string, td *TransferData) error {
	sender, err := users.FindUserByName(senderName)
	if err != nil {
		return err
	}
	recipient, err := users.FindUserByName(senderName)
	if err != nil {
		return err
	}
	if sender.UserID == recipient.UserID {
		return errors.New("can't translate to yourself")
	}
	senderAcc, err := sender.FindAccount(td.AccIDFrom)
	if err != nil {
		return err
	}
	recipientAcc, err := recipient.FindAccount(td.AccIDTo)
	if err != nil {
		return err
	}
	if senderAcc.Balance.Currency != recipientAcc.Balance.Currency {
		return errors.New("different currencies")
	}
	err = senderAcc.TakeOff(td.Balance)
	if err != nil {
		return err
	}
	err = recipientAcc.TopUp(td.Balance)
	if err != nil {
		return err
	}
	return nil
}
