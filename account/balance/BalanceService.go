package balance

import "errors"

func (b *Balance) TopUp(amount float64) {
	b.Amount += amount
}

func (b *Balance) TakeOff(amount float64) error {
	if b.Amount-amount < 0 {
		return errors.New("there are not enough funds in the account")
	}
	b.Amount -= amount
	return nil
}

func (b *Balance) GetCurrency() string {
	switch b.Currency {
	case RUB:
		return "RUB"
	case EUR:
		return "EUR"
	case USD:
		return "USD"
	default:
		return ""
	}
}

func CheckCurrency(currency Currency) error {
	if currency == RUB || currency == EUR || currency == USD {
		return nil
	}
	return errors.New("incorrect currency")
}
