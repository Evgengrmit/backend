package account

import (
	"backend/account/balance"
)

type Account struct {
	ID      uint64          `json:"id"`
	UserID  uint64          `json:"user"`
	Balance balance.Balance `json:"balance"`
}

// Создать аккаунт
// Найти аккаунт по айди
// Получить все аккаунты
// Пополнить счет аккаунта
// Закрыть счет/удалить
