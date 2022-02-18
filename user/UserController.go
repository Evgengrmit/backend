package user

import (
	"backend/account"
	"backend/account/balance"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateUser Создание и сохранения нового пользователя
func CreateUser(c *gin.Context) {
	var creatingUser CreatingUser

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

//func GetUser(c *gin.Context) {
//	name := c.Param("name")
//	foundUser, err := FindUserByName(name)
//	if err != nil {
//		NewErrorResponse(c, http.StatusNotFound, err.Error())
//		return
//	}
//	c.JSON(http.StatusOK, foundUser)
//}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, FindUsers())
}

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

// CreateAccForUser Создание нового аккаунта для юзера
func CreateAccForUser(c *gin.Context) {
	name := c.Param("name")
	var currency balance.Currency
	if err := c.BindJSON(&currency); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	foundUser, err := CreateAccountForUser(name, currency)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, foundUser)
}
