package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
	Balance `json:"balance"`
}

type Currency int64

const (
	RUB Currency = iota
	EUR
	USD
)

type Balance struct {
	Amount float64
	Currency Currency // enum iota
}

var users []User

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/user/{name}", getUser).Methods("GET")
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/user", saveUser).Methods("POST")
	http.Handle("/", router)

	fmt.Println("Server available on port 8080")
	http.ListenAndServe(":8080", nil)
}

func getUsers(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	for _, user := range users {
		if user.Name == name {
			json.NewEncoder(w).Encode(user)
			break
		}
	}
}

func saveUser(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	users = append(users, user)
	json.NewEncoder(w).Encode(&user)
}
