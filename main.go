package main

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type User struct {
	Name    string `json:"name"`
	Age     int32  `json:"age"`
	Balance `json:"balance"`
}

type Currency int64

const (
	NONE Currency = iota
	RUB
	EUR
	USD
)

type Balance struct {
	Amount   float64
	Currency Currency // enum iota
}

var users []User

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/user/{name}", getUser).Methods("GET")
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/user", saveUser).Methods("POST")
	router.HandleFunc("/wallet/top_up", topUpAccount).Methods("POST")

	http.Handle("/", router)

	log.Println("Server available on port 8080")
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

func getUsers(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Println(err)
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	isFind := false
	for _, user := range users {
		if user.Name == name {
			_ = json.NewEncoder(w).Encode(user)
			isFind = true
			break

		}
	}
	if !isFind {
		_, _ = w.Write([]byte("user not found"))
		log.Println("user not found")
	}
}

func saveUser(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	users = append(users, user)
	_ = json.NewEncoder(w).Encode(&user)

}

type Money struct {
	Name     string  `json:"name"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

func checkCurrency(currency string) (Currency, error) {
	switch currency {
	case "RUB":
		return RUB, nil
	case "EUR":
		return EUR, nil
	case "USD":
		return USD, nil
	default:
		return NONE, errors.New("error in the name of the currency")

	}
}
func topUpAccount(w http.ResponseWriter, r *http.Request) {

	var replenishment Money
	_ = json.NewDecoder(r.Body).Decode(&replenishment)
	isFind := false
	for i, user := range users {
		if replenishment.Name == user.Name {
			users[i].Amount += replenishment.Amount
			cur, err := checkCurrency(replenishment.Currency)
			if err != nil {
				_, _ = w.Write([]byte("error in the name of the currency"))
				log.Println("error in the name of the currency")
				return
			}
			users[i].Currency = cur
			_ = json.NewEncoder(w).Encode(&users[i])
			isFind = true
			break
		}

	}
	if !isFind {
		_, _ = w.Write([]byte("it is impossible to top up the account because the user does not exist"))
		log.Println("it is impossible to top up the account because the user does not exist")
	}
}
