package user

import "backend/account/balance"

type TransferData struct {
	AccIDFrom uint64          `json:"id_from"`
	NameTo    string          `json:"name_to"`
	AccIDTo   uint64          `json:"id_to"`
	Balance   balance.Balance `json:"balance"`
}
