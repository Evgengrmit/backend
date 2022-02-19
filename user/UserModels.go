package user

type CreateData struct {
	Name     string `json:"name"`
	Age      int    `json:"age,omitempty"`
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
type UpdateData struct {
	AccountID int     `json:"id"`
	Amount    float64 `json:"amount"`
}
type TransferData struct {
	AccIDFrom int     `json:"id_from"`
	LoginTo   string  `json:"login_to"`
	AccIDTo   int     `json:"id_to"`
	Amount    float64 `json:"amount"`
}

type errorResponse struct {
	Message string `json:"message"`
}
