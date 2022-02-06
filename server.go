package main

import (
	"backend/user"
	"github.com/gin-gonic/gin"
	"os"
)

func GinNewServer() *gin.Engine {
	router := gin.Default()
	router.GET("/users", user.GetUsers)
	router.POST("/user", user.SaveUser)
	router.GET("/user/:name", user.GetUser)
	router.GET("/user/login", user.Authentication)
	router.GET("/user/:name/wallet", user.GetAccountsByName)
	router.PUT("user/:name/wallet/top_up", user.TopUpAccount)
	router.PUT("user/:name/wallet/take_off", user.TakeOffAccount)
	router.PUT("user/:name/wallet/transfer", user.Transfer)
	router.POST("user/:name/wallet/create", user.CreateAccForUser)

	return router
}

func RunServer(port string) error {
	err := os.Setenv("PORT", port)
	if err != nil {
		return err
	}
	router := GinNewServer()
	return router.Run(port)
}
