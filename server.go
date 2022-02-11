package main

import (
	"backend/user"
	"github.com/gin-gonic/gin"
)

func GinNewServer() *gin.Engine {
	router := gin.Default()
	router.GET("/users", user.GetUsers)
	userRouter := router.Group("/user")
	{
		userRouter.POST("", user.CreateUser)
		userRouter.GET("/:name", user.GetUser)
		userRouter.GET("/login", user.Authentication)

		wallet := userRouter.Group("/user/:name/wallet")
		{
			wallet.GET("", user.GetAccountsByName)
			wallet.PUT("/top_up", user.TopUpAccount)
			wallet.PUT("/take_off", user.TakeOffAccount)
			wallet.PUT("/transfer", user.Transfer)
			wallet.POST("/create", user.CreateAccForUser)
		}
	}
	return router
}

func RunServer(port string) error {
	router := GinNewServer()
	return router.Run(port)
}
