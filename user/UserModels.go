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
type UpdateData struct {
	AccountID int64   `json:"id"`
	Amount    float64 `json:"amount"`
}
type TransferData struct {
	AccIDFrom int64   `json:"id_from"`
	LoginTo   string  `json:"name_to"`
	AccIDTo   int64   `json:"id_to"`
	Amount    float64 `json:"amount"`
}

type errorResponse struct {
	Message string `json:"message"`
}
