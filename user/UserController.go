package user

import (
	"backend/account"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateUser Создание и сохранения нового пользователя
func CreateUser(c *gin.Context) {
	var creatingUser CreateData

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
func Deletion(c *gin.Context) {
	var deleteUser LoginData
	if err := c.BindJSON(&deleteUser); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := DeleteUser(deleteUser)
	if err != nil {
		NewErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deletion was successful"})
}

func CreateAccountForUser(c *gin.Context) {
	login := c.Param("login")
	u, err := FindUserByLogin(login)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	var currency account.Currency
	if err := c.BindJSON(&currency); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := u.CreateAccount(currency)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, id)
}

func TopUpAccount(c *gin.Context) {
	login := c.Param("login")
	var uD UpdateData
	foundUser, err := FindUserByLogin(login)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	if err := c.BindJSON(&uD); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = foundUser.TopUpAccount(uD.AccountID, uD.Amount)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func TakeOffAccount(c *gin.Context) {
	login := c.Param("login")
	foundUser, err := FindUserByLogin(login)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	var uD UpdateData
	if err := c.BindJSON(&uD); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = foundUser.TakeOffAccount(uD.AccountID, uD.Amount)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
func Transfer(c *gin.Context) {
	login := c.Param("login")
	foundUser, err := FindUserByLogin(login)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	var transferData TransferData
	if err := c.BindJSON(&transferData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = foundUser.TransferToUserByLogin(transferData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "transfer was successful"})
}
