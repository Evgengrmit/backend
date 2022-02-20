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
	if err := row.Scan(&u.ID, &u.Name, &u.Age, &u.Login, &u.Email, &u.HashedPassword); err != nil {
		return &User{}, err
	}
	err := bcrypt.CompareHashAndPassword(u.HashedPassword, []byte(ld.Password))
	if err != nil {
		return &User{}, err
	}
	return &u, nil
}

func AddNewUser(cU *CreateData) (int64, error) {
	if IsUserExistByLogin(cU.Login) {
		return 0, errors.New("user with this login exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cU.Password), 1)
	var id int64
	err = db.DB.QueryRow("insert into \"user\" (name, age, email, login, password) values ($1, $2, $3, $4, $5) returning id",
		cU.Name, cU.Age, cU.Email, cU.Login, hashedPassword).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func DeleteUser(dd LoginData) error {
	row := db.DB.QueryRow("select id,login,password from \"user\" where login = $1", dd.Login)
	var u User
	if err := row.Scan(&u.ID, &u.Login, &u.HashedPassword); err != nil {
		return err
	}
	err := bcrypt.CompareHashAndPassword(u.HashedPassword, []byte(dd.Password))
	if err != nil {
		return err
	}
	ctx := context.Background()
	tx, err := db.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	_, err = db.DB.Exec("delete from account where  user_id = $1", u.ID)
	if err != nil {
		return tx.Rollback()
	}
	_, err = db.DB.Exec("delete from \"user\" where id = $1", u.ID)
	if err != nil {
		return tx.Rollback()
	}
	return tx.Commit()
}

//
//func FindUserByLogin(login string) (*User, error) {
//	var u User
//	row := db.DB.QueryRow("select id from \"user\" where login = $1", login)
//	err := row.Scan(&u.ID)
//	if err != nil {
//		return nil, err
//	}
//	return &u, nil
//}

func IsUserExistByLogin(login string) bool {
	var exists bool
	err := db.DB.QueryRow("select exists(select * from \"user\" where login = $1)", login).Scan(&exists)
	if err != nil {
		log.Print(err.Error())
	}
	return exists
}

func IsUserExist(id int) bool {
	var exists bool
	err := db.DB.QueryRow("select exists(select * from \"user\" where id = $1)", id).Scan(&exists)
	if err != nil {
		log.Print(err.Error())
	}
	return exists
}

func (u *User) GetAccounts() ([]account.Account, error) {
	return account.FindAccounts(u.ID)
}

func TopUpAccount(id int, amount float64) error {
	foundAccount, err := account.FindAccount(id)
	if err != nil {
		return err
	}
	err = foundAccount.TopUp(amount)
	if err != nil {
		return err
	}
	return nil

}

func TakeOffAccount(id int, amount float64) error {
	foundAccount, err := account.FindAccount(id)
	if err != nil {
		return err
	}
	err = foundAccount.TakeOff(amount)
	if err != nil {
		return err
	}
	return nil
}

func TransferToUserByLogin(td TransferData) error {
	accountFrom, err := account.FindAccount(td.AccIDFrom)
	if err != nil {
		return err
	}
	accountTo, err := account.FindAccount(td.AccIDTo)
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
