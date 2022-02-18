package user

type CreateUserData struct {
	Name     string `json:"name"`
	Age      int8   `json:"age,omitempty"`
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type TransferData struct {
	AccIDFrom uint64 `json:"id_from"`
	NameTo    string `json:"name_to"`
	AccIDTo   uint64 `json:"id_to"`
	//Balance   balance.Balance `json:"balance"`
}

type errorResponse struct {
	Message string `json:"message"`
}
