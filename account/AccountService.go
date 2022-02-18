package account

import (
	"backend/db"
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

//
//func (a *Account) TopUp(b balance.Balance) error {
//	err := balance.CheckCurrency(b.Currency)
//	if err != nil {
//		return err
//	}
//	if a.Balance.Currency == b.Currency {
//		// Обновление поле amount в таблице аккаунтов
//		a.Balance.TopUp(b.Amount)
//		return nil
//	}
//	return fmt.Errorf("different currencies %d %d", a.Balance.Currency, b.Currency)
//}
//
//func (a *Account) TakeOff(b balance.Balance) error {
//	err := balance.CheckCurrency(b.Currency)
//	if err != nil {
//		return err
//	}
//	if a.Balance.Currency == b.Currency {
//		err = a.Balance.TakeOff(b.Amount)
//		if err != nil {
//			return err
//		}
//		return nil
//	}
//	return errors.New("different currencies")
//}
