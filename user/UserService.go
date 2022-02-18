package user

import (
	"backend/account"
	"backend/db"
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

func FindUserByLogin(login string) *User {
	var u User
	row := db.DB.QueryRow("select id from \"user\" where login = $1)", login)
	err := row.Scan(&u.ID)
	if err != nil {
		return nil
	}
	return &u
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
	accounts, err := account.FindAccounts(u.ID)
	if err != nil {
		return nil, err
	}
	u.Accounts = accounts
	return u.Accounts, nil
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

func TopUpForUser(username string, accData *account.Account) (User, error) {
	//foundUser, err := users.FindUserByName(username)
	//if err != nil {
	//	return User{}, err
	//}
	//err = foundUser.TopUpAccount(accData.AccountID, accData.Balance)
	//if err != nil {
	//	return User{}, err
	//}
	//return *foundUser, nil
	return User{}, nil
}
func TakeOffForUser(username string, accData *account.Account) (User, error) {
	//foundUser, err := users.FindUserByName(username)
	//if err != nil {
	//	return User{}, err
	//}
	//err = foundUser.TakeOffAccount(accData.AccountID, accData.Balance)
	//if err != nil {
	//	return User{}, err
	//}
	return User{}, nil
}
func TransferBetweenUsers(senderName string, td *TransferData) error {
	//sender, err := users.FindUserByName(senderName)
	//if err != nil {
	//	return err
	//}
	//recipient, err := users.FindUserByName(senderName)
	//if err != nil {
	//	return err
	//}
	//if sender.ID == recipient.ID {
	//	return errors.New("can't translate to yourself")
	//}
	//senderAcc, err := sender.FindAccount(td.AccIDFrom)
	//if err != nil {
	//	return err
	//}
	//recipientAcc, err := recipient.FindAccount(td.AccIDTo)
	//if err != nil {
	//	return err
	//}
	//if senderAcc.Balance.Currency != recipientAcc.Balance.Currency {
	//	return errors.New("different currencies")
	//}
	//err = senderAcc.TakeOff(td.Balance)
	//if err != nil {
	//	return err
	//}
	//err = recipientAcc.TopUp(td.Balance)
	//if err != nil {
	//	return err
	//}
	return nil
}
