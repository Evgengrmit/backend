package user

import (
	"backend/account"
	"backend/account/balance"
	"backend/db"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func NewUser() *User {
	return &User{ID: uint64(users.NumberOfUsers() + 1)}
}

func (u *User) FindAccount(id uint64) (*account.Account, error) {
	for i := range u.Accounts {
		if id != u.Accounts[i].AccountID {
			return &u.Accounts[i], nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Account %d not found", id))
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

func (u *User) CreateAccount(currency balance.Currency) error {
	err := balance.CheckCurrency(currency)
	if err != nil {
		return err
	}
	newAccount := account.Account{
		AccountID: uint64(len(accounts) + 1),
		UserID:    u.ID,
		Balance:   balance.Balance{Currency: currency}}
	u.Accounts = append(u.Accounts, newAccount)
	accounts = append(accounts, newAccount)
	return nil
}

// USERS

func (users *Users) LoginUser(ld *LoginData) (*User, error) {
	foundUser, err := users.FindUserByLogin(ld.Login)
	if err != nil {
		return &User{}, err
	}
	err = bcrypt.CompareHashAndPassword(foundUser.HashedPassword, []byte(ld.Password))
	if err != nil {
		return &User{}, err
	}
	return foundUser, nil

}
func (users *Users) CreatingUser(cU *CreatingUser) (User, error) {
	u := NewUser()
	u.Name = cU.Name
	u.Age = cU.Age
	u.Email = cU.Email
	fmt.Println("Test5")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cU.Password), 1)
	fmt.Println("Test6")
	if err != nil {
		return User{}, err
	}
	u.HashedPassword = hashedPassword
	fmt.Println("Start sending insert")
	_, err = db.DB.Exec("insert into \"user\" (name, age, email, login, password) values ($1, $2, $3, $4, $5)",
		u.Name, u.Age, u.Email, u.Login, u.HashedPassword)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Finish sending insert")
	users.Users = append(users.Users, *u)
	return *u, nil
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

func (users *Users) FindUserByLogin(login string) (*User, error) {
	for i := range users.Users {
		if users.Users[i].Login == login {
			return &users.Users[i], nil
		}
	}
	return nil, errors.New(fmt.Sprintf("User %s not found", login))
}

func (users *Users) GetUsers() []User {
	return users.Users
}

func (users *Users) CreateAccountForUser(username string, curr balance.Currency) (User, error) {
	foundUser, err := users.FindUserByName(username)
	if err != nil {
		return User{}, err
	}
	if err = foundUser.CreateAccount(curr); err != nil {
		return User{}, err
	}
	return *foundUser, nil
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
	if sender.ID == recipient.ID {
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
