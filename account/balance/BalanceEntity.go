package balance

type Balance struct {
	Amount   float64  `json:"amount"`
	Currency Currency `json:"currency"` // enum iota
}

type Currency int64

const (
	RUB Currency = iota
	EUR
	USD
)
