package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type User struct {
	Name    string `json:"name"`
	Age     int32  `json:"age"`
	Balance `json:"balance"`
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
	RUB Currency = iota
	EUR
	USD
)

type Balance struct {
	Amount   float64  `json:"amount"`
	Currency Currency `json:"currency"` // enum iota
}

var users []User

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user", saveUser).Methods("POST")
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/user/{name}", getUser).Methods("GET")
	router.HandleFunc("/user/{name}/wallet", checkTheBalance).Methods("GET")
	router.HandleFunc("/user/{name}/wallet/top_up", topUpAccount).Methods("POST")
	router.HandleFunc("/user/{name}/wallet/take_off", takeOffAccount).Methods("PUT")
	router.HandleFunc("/user/{name}/wallet/transfer", transfer).Methods("POST")
	http.Handle("/", router)
	log.Println("Server available on port 8080")
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

func saveUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		log.Println(fmt.Errorf("function saveUser: %s", err.Error()))
		return
	}
	users = append(users, user)
	err = json.NewEncoder(w).Encode(&user)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		log.Println(fmt.Errorf("function saveUser: %s", err.Error()))
		return
	}
}

func getUsers(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		log.Println(fmt.Errorf("function getUsers: %s", err.Error()))
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	isFind := false
	for _, user := range users {
		if user.Name == name {
			err := json.NewEncoder(w).Encode(user)
			if err != nil {
				_, _ = w.Write([]byte(err.Error()))
				log.Println(fmt.Errorf("function getUser: %s", err.Error()))
				return
			}
			isFind = true
			break

		}
	}
	if !isFind {
		_, _ = w.Write([]byte("user not found"))
		log.Println("function getUser: user not found")
	}
}

func checkTheBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	isFind := false
	for _, user := range users {
		if user.Name == name {
			balance := &Balance{Amount: user.Amount, Currency: user.Currency}
			err := json.NewEncoder(w).Encode(balance)
			if err != nil {
				_, _ = w.Write([]byte(err.Error()))
				log.Println(fmt.Errorf("function checkTheBalance: %s", err.Error()))
				return
			}
			isFind = true
			break
		}
	}
	if !isFind {
		_, _ = w.Write([]byte("user not found"))
		log.Println("function checkTheBalance: user not found")
	}

}

func topUpAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	var replenishment Balance
	err := json.NewDecoder(r.Body).Decode(&replenishment)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		log.Println(fmt.Errorf("function topUpAccount: %s", err.Error()))
		return
	}
	isFind := false
	for i, user := range users {
		if name == user.Name {
			users[i].Amount += replenishment.Amount
			err = checkCurrency(replenishment.Currency)
			if err != nil {
				_, _ = w.Write([]byte(err.Error()))
				log.Println(fmt.Errorf("function topUpAccount: %s", err.Error()))
				return
			}
			users[i].Currency = replenishment.Currency
			err = json.NewEncoder(w).Encode(&users[i])
			if err != nil {
				_, _ = w.Write([]byte(err.Error()))
				log.Println(fmt.Errorf("function topUpAccount: %s", err.Error()))
				return
			}
			isFind = true
			break
		}

	}
	if !isFind {
		_, _ = w.Write([]byte("it is impossible to top up the account because the user does not exist"))
		log.Println("function topUpAccount: it is impossible to top up the account because the user does not exist")
	}
}

func takeOffAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	var withdrawal Balance
	err := json.NewDecoder(r.Body).Decode(&withdrawal)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		log.Println(fmt.Errorf("function takeOffAccount: %s", err.Error()))
		return
	}
	isFind := false
	for i, user := range users {
		if name == user.Name {
			err := checkCurrency(withdrawal.Currency)
			if err != nil {
				_, _ = w.Write([]byte(err.Error()))
				log.Println(fmt.Errorf("function takeOffAccount: %s", err.Error()))
				return
			}
			if user.Currency != withdrawal.Currency {
				_, _ = w.Write([]byte(fmt.Sprintf("other currency %s\n", user.getCurrency())))
				log.Printf("function takeOffAccount: other currency %s\n", user.getCurrency())
				return
			}
			if user.Amount-withdrawal.Amount >= 0 {
				users[i].Amount -= withdrawal.Amount
				err = json.NewEncoder(w).Encode(&users[i])
				if err != nil {
					_, _ = w.Write([]byte(err.Error()))
					log.Println(fmt.Errorf("function takeOffAccount: %s", err.Error()))
					return
				}
				isFind = true
				break
			} else {
				_, _ = w.Write([]byte("there are not enough funds in the account"))
				log.Println("function takeOffAccount: there are not enough funds in the account")
				return
			}
		}
	}
	if !isFind {
		_, _ = w.Write([]byte("it is impossible to take off the account because the user does not exist"))
		log.Println("function takeOffAccount: it is impossible to take off the account because the user does not exist")
	}
}

func transfer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	sender, recipient, transferData := &User{}, &User{}, &User{}
	isFindSender, isFindRecipient := false, false
	err := json.NewDecoder(r.Body).Decode(transferData)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		log.Println(fmt.Errorf("function transfer: %s", err.Error()))
		return
	}
	for i, user := range users {
		if user.Name == name {
			sender = &users[i]
			isFindSender = true

		}
		if user.Name == transferData.Name {
			recipient = &users[i]
			isFindRecipient = true
		}
		if isFindSender && isFindRecipient {
			if sender.Amount-transferData.Amount >= 0 {
				sender.Amount -= transferData.Amount
				if sender.Currency == RUB {
					recipient.Amount += transferData.Amount * 100
				} else {
					recipient.Amount += transferData.Amount
				}
				_, _ = w.Write([]byte("the transfer was carried out successfully"))
			} else {
				_, _ = w.Write([]byte("there are not enough funds in the account"))
				log.Println("function transfer: there are not enough funds in the account")
			}
			break
		}
	}
	if !(isFindSender && isFindRecipient) {
		_, _ = w.Write([]byte("user not found"))
		log.Println("function transfer: user not found")
	}

}
