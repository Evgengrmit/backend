package main

import (
	"backend/user"
	"github.com/gin-gonic/gin"
)

func GinNewServer() *gin.Engine {
	router := gin.Default()
	userRouter := router.Group("/user")
	{
		userRouter.POST("", user.CreateUser)
		userRouter.GET("/login", user.Authentication)

		wallet := userRouter.Group("/user/:login/wallet")
		{
			//wallet.GET("", user.GetAccountsByName)
			wallet.POST("/create", user.CreateAccountForUser)
			wallet.PUT("/top_up", user.TopUpAccount)
			wallet.PUT("/take_off", user.TakeOffAccount)
			wallet.PUT("/transfer", user.Transfer)

		}
	}
	return router
}

func RunServer(port string) error {
	router := GinNewServer()
	return router.Run(port)
}
