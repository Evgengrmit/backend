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

	if err := c.ShouldBindJSON(&creatingUser); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := users.CreatingUser(&creatingUser)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, user)
}

func Authentication(c *gin.Context) {
	var loginUser LoginData
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := users.LoginUser(&loginUser)
	if err != nil {
		NewErrorResponse(c, http.StatusForbidden, err.Error())
	}
	c.JSON(http.StatusOK, user)
}

func GetUser(c *gin.Context) {
	name := c.Param("name")
	foundUser, err := users.FindUserByName(name)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, foundUser)
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users.GetUsers())
}

func GetAccountsByName(c *gin.Context) {
	name := c.Param("name")
	foundUser, err := users.FindUserByName(name)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, foundUser.GetAccounts())
}

// TopUpAccount Пополнение счета
func TopUpAccount(c *gin.Context) {
	name := c.Param("name")
	var replenishment account.Account
	if err := c.ShouldBindJSON(&replenishment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	foundUser, err := users.TopUpForUser(name, &replenishment)
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
	if err := c.ShouldBindJSON(&replenishment); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	foundUser, err := users.TakeOffForUser(name, &replenishment)
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
	if err := c.ShouldBindJSON(&transferData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := users.TransferBetweenUsers(name, &transferData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"status": "transfer was successful"})
}

// CreateAccForUser Создание нового аккаунта для юзера
func CreateAccForUser(c *gin.Context) {
	name := c.Param("name")
	var currency balance.Currency
	if err := c.ShouldBindJSON(&currency); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	foundUser, err := users.CreateAccountForUser(name, currency)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, foundUser)
}
