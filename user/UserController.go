package user

import (
	"backend/account"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SignUp Создание и сохранение нового пользователя
func SignUp(c *gin.Context) {
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
	c.JSON(http.StatusCreated, &gin.H{"id": userID})
}

func SignIn(c *gin.Context) {
	var loginUser LoginData
	if err := c.BindJSON(&loginUser); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userToken, err := GenerateToken(&loginUser)
	if err != nil {
		NewErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	c.JSON(http.StatusOK, &gin.H{"token": userToken})
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
	userID, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	var currency account.Currency
	if err := c.BindJSON(&currency); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := account.AddNewAccount(userID, currency)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, id)
}

func TopUpAccountForUser(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	var uD UpdateData
	if err := c.BindJSON(&uD); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = TopUpAccount(uD.AccountID, userID, uD.Amount)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func TakeOffAccountForUser(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	var uD UpdateData
	if err := c.BindJSON(&uD); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = TakeOffAccount(uD.AccountID, userID, uD.Amount)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
func Transfer(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	var transferData TransferData
	if err := c.BindJSON(&transferData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = TransferToUserByLogin(userID, transferData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "transfer was successful"})
}
