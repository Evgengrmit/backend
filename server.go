package main

import (
	"backend/authorization"
	"backend/user"
	"github.com/gin-gonic/gin"
)

func GinNewServer() *gin.Engine {
	router := gin.Default()
	auth := router.Group("/auth")
	{
		auth.POST("sign_up", authorization.SignUp)
		auth.POST("sign_in", authorization.SignIn)
	}
	userRouter := router.Group("/user", user.Identity)
	{
		userRouter.DELETE("/delete", user.Deletion)

		wallet := userRouter.Group("/accounts")
		{
			//wallet.GET("", user.GetAccountsByName)
			wallet.POST("/create", user.CreateAccountForUser)
			wallet.PUT("/top_up", user.TopUpAccountForUser)
			wallet.PUT("/take_off", user.TakeOffAccountForUser)
			wallet.PUT("/transfer", user.Transfer)

		}
	}
	return router
}

func RunServer(port string) error {
	router := GinNewServer()
	return router.Run(port)
}
