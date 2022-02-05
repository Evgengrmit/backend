package main

import (
	"backend/user"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func NewServer() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users", user.GetUsers).Methods("GET")

	userRouter := router.PathPrefix("/user").Subrouter()

	userRouter.HandleFunc("", user.SaveUser).Methods("POST")
	userRouter.HandleFunc("/{name}", user.GetUser).Methods("GET")
	userRouter.HandleFunc("/{name}/wallet", user.GetAccountsByName).Methods("GET")
	userRouter.HandleFunc("/{name}/wallet/top_up", user.TopUpAccount).Methods("POST")
	userRouter.HandleFunc("/{name}/wallet/take_off", user.TakeOffAccount).Methods("PUT")
	userRouter.HandleFunc("/{name}/wallet/transfer", user.Transfer).Methods("POST")

	return router
}

func RunServer(port string) error {
	os.Setenv("PORT", port)
	router := NewServer()
	http.Handle("/", router)
	return http.ListenAndServe(":"+port, nil)
}
