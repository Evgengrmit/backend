package user

import (
	"backend/account"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateUser Создание и сохранения нового пользователя
func CreateUser(c *gin.Context) {
	var creatingUser CreateUserData

	if err := c.BindJSON(&creatingUser); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userID, err := AddNewUser(&creatingUser)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, userID)
}

func Authentication(c *gin.Context) {
	var loginUser LoginData
	if err := c.BindJSON(&loginUser); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := LoginUser(&loginUser)
	if err != nil {
		NewErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateAccountForUser(c *gin.Context) {
	login := c.Param("login")
	var currency account.Currency
	if err := c.BindJSON(&currency); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := FindUserByLogin(login)
	if err := u.CreateAccount(currency); err != nil {
		NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

//func GetUser(c *gin.Context) {
//	name := c.Param("name")
//	foundUser, err := FindUserByName(name)
//	if err != nil {
//		NewErrorResponse(c, http.StatusNotFound, err.Error())
//		return
//	}
//	c.JSON(http.StatusOK, foundUser)
//}

//func GetAccountsByName(c *gin.Context) {
//	name := c.Param("name")
//	foundUser, err := FindUserByName(name)
//	if err != nil {
//		NewErrorResponse(c, http.StatusNotFound, err.Error())
//		return
//	}
//	c.JSON(http.StatusOK, foundUser.GetAccounts())
//}

// TopUpAccount Пополнение счета
func TopUpAccount(c *gin.Context) {
	name := c.Param("name")
	var replenishment account.Account
	if err := c.BindJSON(&replenishment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	foundUser, err := TopUpForUser(name, &replenishment)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, foundUser)
}

// TakeOffAccount Пополнение счета
func TakeOffAccount(c *gin.Context) {
	name := c.Param("name")
	var replenishment account.Account
	if err := c.BindJSON(&replenishment); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	foundUser, err := TakeOffForUser(name, &replenishment)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, foundUser)
}

// Transfer Перевод между пользователями
func Transfer(c *gin.Context) {
	name := c.Param("name")
	var transferData TransferData
	if err := c.BindJSON(&transferData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := TransferBetweenUsers(name, &transferData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"status": "transfer was successful"})
}

// CreateAccountForUser Создание нового аккаунта для юзера
