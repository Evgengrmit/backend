package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

type User struct {
	Name    string `json:"name"`
	Age     int32  `json:"age"`
	Balance `json:"balance"`
}

func (u *User) topUpBalance(money float64) {
	u.Amount += money
}
func (u *User) takeOffBalance(money float64) {
	u.Amount -= money
}
func (u *User) getCurrency() string {
	switch u.Currency {
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

func checkCurrency(currency Currency) error {
	if currency == RUB || currency == EUR || currency == USD {
		return nil
	}
	return errors.New("incorrect currency")
}

type Currency int64

const (
	coefficient          = 100
	RUB         Currency = iota
	EUR
	USD
)

type Balance struct {
	Amount   float64  `json:"amount"`
	Currency Currency `json:"currency"` // enum iota
}

func translateError(w io.Writer, err error, funcName string) {
	_, _ = w.Write([]byte(err.Error()))
	log.Println(fmt.Errorf("ERROR\tfunction %s: %s", funcName, err.Error()))
}
func translateInfo(w io.Writer, info string, funcName string) {
	_, _ = w.Write([]byte(info))
	log.Println(fmt.Sprintf("INFO\tfunction %s: %s", funcName, info))
}

var users []User

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user", saveUser).Methods("POST")
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/user/{name}", getUser).Methods("GET")
	router.HandleFunc("/user/{name}/wallet", getTheBalance).Methods("GET")
	router.HandleFunc("/user/{name}/wallet/top_up", topUpAccount).Methods("POST")
	router.HandleFunc("/user/{name}/wallet/take_off", takeOffAccount).Methods("PUT")
	router.HandleFunc("/user/{name}/wallet/transfer", transfer).Methods("POST")
	http.Handle("/", router)
	log.Println("Server available on port 8080")
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

// Создание и сохранения нового пользователя
func saveUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		translateError(w, err, "saveUser")
		return
	}
	users = append(users, user)
	err = json.NewEncoder(w).Encode(&user)
	if err != nil {
		translateError(w, err, "saveUser")
		return
	}
}

// Список всех пользователей
func getUsers(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		translateError(w, err, "getUsers")
	}
}

// Данные о пользователе
func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	for _, user := range users {
		if user.Name == name {
			err := json.NewEncoder(w).Encode(user)
			if err != nil {
				translateError(w, err, "getUser")
				return
			}

		}
	}
	err := errors.New("user not found")
	translateError(w, err, "getUser")

}

// Баланс пользователя
func getTheBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	for _, user := range users {
		if user.Name == name {
			balance := &Balance{Amount: user.Amount, Currency: user.Currency}
			err := json.NewEncoder(w).Encode(balance)
			if err != nil {
				translateError(w, err, "getTheBalance")
				return
			}
		}
	}
	err := errors.New("user not found")
	translateError(w, err, "getTheBalance")

}

// Пополнение счета
func topUpAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	var replenishment Balance
	err := json.NewDecoder(r.Body).Decode(&replenishment)
	if err != nil {
		translateError(w, err, "topUpAccount")
		return
	}

	for i := range users {
		if name == users[i].Name {
			err = checkCurrency(replenishment.Currency)
			if err != nil {
				translateError(w, err, "topUpAccount")
				return
			}
			users[i].topUpBalance(replenishment.Amount)
			users[i].Currency = replenishment.Currency
			err = json.NewEncoder(w).Encode(users[i])
			if err != nil {
				translateError(w, err, "topUpAccount")
			}
			return
		}

	}
	err = errors.New("it is impossible to top up the account because the user does not exist")
	translateError(w, err, "topUpAccount")
}

//Снятие со счета
func takeOffAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	var withdrawal Balance
	err := json.NewDecoder(r.Body).Decode(&withdrawal)
	if err != nil {
		translateError(w, err, "takeOffAccount")
		return
	}
	for i := range users {
		if name == users[i].Name {
			err = checkCurrency(withdrawal.Currency)
			if err != nil {
				translateError(w, err, "takeOffAccount")
				return
			}
			if users[i].Currency != withdrawal.Currency {
				err = fmt.Errorf("other currency %s\n", users[i].getCurrency())
				translateError(w, err, "takeOffAccount")
				return
			}
			if users[i].Amount-withdrawal.Amount >= 0 {
				users[i].takeOffBalance(withdrawal.Amount)
				err = json.NewEncoder(w).Encode(users[i])
				if err != nil {
					translateError(w, err, "takeOffAccount")
					return
				}
			} else {
				err = errors.New("there are not enough funds in the account")
				translateError(w, err, "takeOffAccount")
				return
			}
		}
	}
	err = errors.New("it is impossible to take off the account because the user does not exist")
	translateError(w, err, "takeOffAccount")
}

//Перевод между пользователями
func transfer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	sender, recipient, transferData := &User{}, &User{}, &User{}
	isFindSender, isFindRecipient := false, false
	err := json.NewDecoder(r.Body).Decode(transferData)
	if err != nil {
		translateError(w, err, "transfer")
		return
	}
	for i := range users {
		if users[i].Name == name {
			sender = &users[i]
			isFindSender = true

		}
		if users[i].Name == transferData.Name {
			recipient = &users[i]
			isFindRecipient = true
		}
		if isFindSender && isFindRecipient {
			if sender.Name == recipient.Name {
				err = errors.New("can't translate to yourself")
				translateError(w, err, "transfer")
				return
			}
			if sender.Amount-transferData.Amount >= 0 {
				sender.takeOffBalance(transferData.Amount)
				if sender.Currency == RUB {
					recipient.topUpBalance(transferData.Amount * coefficient)
				} else {
					recipient.topUpBalance(transferData.Amount)
				}
				translateInfo(w, "the transfer was carried out successfully", "transfer")

			} else {
				err = errors.New("there are not enough funds in the account")
				translateError(w, err, "transfer")
			}
			return
		}
	}
	err = errors.New("user not found")
	translateError(w, err, "transfer")

}
