package account

type Account struct {
	ID       int64    `json:"id"`
	UserID   int64    `json:"user"`
	Amount   float64  `json:"amount"`
	Currency Currency `json:"currency"` // enum iota
}
type Currency int64

const (
	RUB Currency = iota
	EUR
	USD
)
