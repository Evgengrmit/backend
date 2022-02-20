package user

import (
	"backend/account"
	"backend/myerr"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SignUp Создание и сохранение нового пользователя

func Deletion(c *gin.Context) {
	var deleteUser LoginData
	if err := c.BindJSON(&deleteUser); err != nil {
		myerr.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := DeleteUser(deleteUser)
	if err != nil {
		myerr.NewErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deletion was successful"})
}

func CreateAccountForUser(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		myerr.NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	var currency account.Currency
	if err := c.BindJSON(&currency); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := account.AddNewAccount(userID, currency)
	if err != nil {
		myerr.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, id)
}

func TopUpAccountForUser(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		myerr.NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	var uD UpdateData
	if err := c.BindJSON(&uD); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = TopUpAccount(uD.AccountID, uD.Amount)
	if err != nil {
		myerr.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func TakeOffAccountForUser(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		myerr.NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	var uD UpdateData
	if err := c.BindJSON(&uD); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = TakeOffAccount(uD.AccountID, uD.Amount)
	if err != nil {
		myerr.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
func Transfer(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		myerr.NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	var transferData TransferData
	if err := c.BindJSON(&transferData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = TransferToUserByLogin(transferData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "transfer was successful"})
}
