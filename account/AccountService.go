package account

import (
	"backend/db"
	"errors"
)

func AddNewAccount(user_id int64, currency Currency) (int64, error) {
	ex, err := db.DB.Exec("insert into \"account\" (currency, user_id) values ($1,$2)", currency, user_id)
	if err != nil {
		return 0, err
	}
	return ex.LastInsertId()
}
func FindAccount(accountId, userId int64) (*Account, error) {
	row := db.DB.QueryRow("select * from \"account\" where id = $1 and user_id = $2", accountId, userId)
	var acc Account
	if err := row.Scan(&acc.ID, &acc.Amount, &acc.Currency, &acc.UserID); err != nil {
		return &Account{}, err
	}

	return &acc, nil
}

func FindAccounts(userId int64) ([]Account, error) {
	rows, err := db.DB.Query("select * from \"account\" where user_id = $1", userId)
	if err != nil {
		return nil, err
	}
	var results []Account
	for rows.Next() {
		acc := Account{}
		err = rows.Scan(&acc.ID, &acc.Amount, &acc.Currency, &acc.UserID)
		if err != nil {
			return nil, err
		}
		results = append(results, acc)
	}
	return results, nil
}

func (a *Account) TopUp(amount float64) error {
	_, err := db.DB.Exec("update \"account\" set amount = amount + $1 where id=$2", amount, a.ID)
	return err
}

func (a *Account) TakeOff(amount float64) error {
	row := db.DB.QueryRow("select amount from account where id=$1", a.ID)
	var currentBalance float64
	err := row.Scan(&currentBalance)
	if err != nil {
		return err
	}
	if currentBalance-amount > 0 {
		_, err := db.DB.Exec("update \"account\" set amount = amount - $1 where id=$2", amount, a.ID)
		return err
	}
	return errors.New("insufficient funds to withdraw")
}
