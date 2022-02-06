package user

import (
	"backend/account"
	"backend/account/balance"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SaveUser Создание и сохранения нового пользователя
func SaveUser(c *gin.Context) {
	user := NewUser()

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users.AddUser(*user)
	c.JSON(http.StatusCreated, user)
}

func GetUser(c *gin.Context) {
	name := c.Param("name")
	foundUser, err := users.FindUserByName(name)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
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
		c.String(http.StatusNotFound, err.Error())
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
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, foundUser)
}

// TakeOffAccount Пополнение счета
func TakeOffAccount(c *gin.Context) {
	name := c.Param("name")
	var replenishment account.Account
	if err := c.ShouldBindJSON(&replenishment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	foundUser, err := users.TakeOffForUser(name, &replenishment)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
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
	c.String(http.StatusOK, "transfer was successful")
}

// CreateAccForUser Создание нового аккаунта для юзера
func CreateAccForUser(c *gin.Context) {
	name := c.Param("name")
	var currency balance.Currency
	if err := c.ShouldBindJSON(&currency); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	foundUser, err := users.CreateAccountForUser(name, currency)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, foundUser)
}
