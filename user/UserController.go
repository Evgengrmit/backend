package user

import (
	"backend/account"
	"backend/helper"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// SaveUser Создание и сохранения нового пользователя
func SaveUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := NewUser()
	err := json.NewDecoder(r.Body).Decode(&user)
	users.AddUser(*user)
	helper.CheckError(w, err, "saveUser")
	err = json.NewEncoder(w).Encode(&user)
	if err != nil {
		helper.TranslateError(w, err, "saveUser")
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	name := vars["name"]
	foundUser, err := users.FindUserByName(name)
	if err != nil {
		helper.TranslateError(w, err, "getUser")
		return
	}
	err = json.NewEncoder(w).Encode(foundUser)
	if err != nil {
		helper.TranslateError(w, err, "getUser")
		return
	}
}

func GetUsers(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allUsers := users.GetUsers()
	err := json.NewEncoder(w).Encode(allUsers)
	if err != nil {
		helper.TranslateError(w, err, "getUsers")
	}
}

func GetAccountsByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	name := vars["name"]
	foundUser, err := users.FindUserByName(name)
	if err != nil {
		helper.TranslateError(w, err, "GetBalancesByName")
		return
	}
	err = json.NewEncoder(w).Encode(foundUser.GetAccounts())
	if err != nil {
		helper.TranslateError(w, err, "getBalanceByName")
		return
	}
}

// TopUpAccount Пополнение счета
func TopUpAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	var replenishment account.Account
	err := json.NewDecoder(r.Body).Decode(&replenishment)
	if err != nil {
		helper.TranslateError(w, err, "topUpAccount")
		return
	}
	foundUser, err := users.TakeOffForUser(name, &replenishment)
	if err != nil {
		helper.TranslateError(w, err, "TopUpAccount")
		return
	}
	err = json.NewEncoder(w).Encode(foundUser)
	if err != nil {
		helper.TranslateError(w, err, "TopUpAccount")
		return
	}
	return
}

// TakeOffAccount Снятие со счета
func TakeOffAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	var replenishment account.Account
	err := json.NewDecoder(r.Body).Decode(&replenishment)
	if err != nil {
		helper.TranslateError(w, err, "takeOffAccount")
		return
	}
	foundUser, err := users.TakeOffForUser(name, &replenishment)
	if err != nil {
		helper.TranslateError(w, err, "takeOffAccount")
		return
	}
	err = json.NewEncoder(w).Encode(foundUser)
	if err != nil {
		helper.TranslateError(w, err, "takeOffAccount")
		return
	}
	return
}

// Transfer Перевод между пользователями
func Transfer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	transferData := &TransferData{}
	err := json.NewDecoder(r.Body).Decode(transferData)
	if err != nil {
		helper.TranslateError(w, err, "Transfer")
		return
	}
	err = users.TransferBetweenUsers(name, transferData)
	if err != nil {
		helper.TranslateError(w, err, "Transfer")
		return
	}
	return
}
