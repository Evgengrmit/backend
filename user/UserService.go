package user

import (
	"backend/account"
	"backend/db"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func LoginUser(ld *LoginData) (*User, error) {
	row := db.DB.QueryRow("select * from \"user\" where login = $1", ld.Login)
	var u User
	if err := row.Scan(&u.ID, &u.Name, &u.Age, &u.Email, &u.Login, &u.HashedPassword); err != nil {
		return &User{}, err
	}
	err := bcrypt.CompareHashAndPassword(u.HashedPassword, []byte(ld.Password))
	if err != nil {
		return &User{}, err
	}
	return &u, nil
}

func AddNewUser(cU *CreateUserData) (int64, error) {
	if IsUserExist(cU.Login) {
		return 0, errors.New("user with this login exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cU.Password), 1)
	ex, err := db.DB.Exec("insert into \"user\" (name, age, email, login, password) values ($1, $2, $3, $4, $5)",
		cU.Name, cU.Age, cU.Email, cU.Login, hashedPassword)
	if err != nil {
		return 0, err
	}
	return ex.LastInsertId()
}

func FindUserByLogin(login string) (*User, error) {
	var u User
	row := db.DB.QueryRow("select id from \"user\" where login = $1", login)
	err := row.Scan(&u.ID)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func IsUserExist(login string) bool {
	var exists bool
	err := db.DB.QueryRow("select exists(select * from \"user\" where login = $1)", login).Scan(&exists)
	if err != nil {
		log.Print(err.Error())
	}
	return exists
}

func (u *User) CreateAccount(currency account.Currency) error {
	_, err := account.AddNewAccount(u.ID, currency)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) FindAccount(id int64) (*account.Account, error) {
	return account.FindAccount(id, u.ID)
}

func (u *User) GetAccounts() ([]account.Account, error) {
	return account.FindAccounts(u.ID)
}

func (u *User) TopUpAccount(id int64, amount float64) error {
	foundAcc, err := account.FindAccount(id, u.ID)
	if err != nil {
		return err
	}
	err = foundAcc.TopUp(amount)
	if err != nil {
		return err
	}
	return nil

}

func (u *User) TakeOffAccount(id int64, amount float64) error {
	foundAcc, err := account.FindAccount(id, u.ID)
	if err != nil {
		return err
	}
	err = foundAcc.TakeOff(amount)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) TransferToUserByLogin(td TransferData) error {
	userTo, err := FindUserByLogin(td.LoginTo)
	if err != nil {
		return err
	}
	accountFrom, err := u.FindAccount(td.AccIDFrom)
	if err != nil {
		return err
	}
	accountTo, err := userTo.FindAccount(td.AccIDTo)
	if err != nil {
		return err
	}
	ctx := context.Background()
	tx, err := db.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	err = accountFrom.TakeOff(td.Amount)
	if err != nil {
		return tx.Rollback()
	}
	err = accountTo.TopUp(td.Amount)
	if err != nil {
		return tx.Rollback()
	}
	return tx.Commit()
}
