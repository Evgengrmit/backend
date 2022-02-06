package account

import (
	"backend/account/balance"
	"errors"
	"fmt"
)

func (a *Account) TopUp(b balance.Balance) error {
	err := balance.CheckCurrency(b.Currency)
	if err != nil {
		return err
	}
	if a.Balance.Currency == b.Currency {
		a.Balance.TopUp(b.Amount)
		return nil
	}
	return fmt.Errorf("different currencies %d %d", a.Balance.Currency, b.Currency)
}

func (a *Account) TakeOff(b balance.Balance) error {
	err := balance.CheckCurrency(b.Currency)
	if err != nil {
		return err
	}
	if a.Balance.Currency == b.Currency {
		err = a.Balance.TakeOff(b.Amount)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("different currencies")
}
