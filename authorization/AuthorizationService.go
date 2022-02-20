package authorization

import (
	"backend/myerr"
	"backend/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUp(c *gin.Context) {
	var creatingUser user.CreateData

	if err := c.BindJSON(&creatingUser); err != nil {
		myerr.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userID, err := user.AddNewUser(&creatingUser)

	if err != nil {
		myerr.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, &gin.H{"id": userID})
}

func SignIn(c *gin.Context) {
	var loginUser user.LoginData
	if err := c.BindJSON(&loginUser); err != nil {
		myerr.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userToken, err := user.GenerateToken(&loginUser)
	if err != nil {
		myerr.NewErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	c.JSON(http.StatusOK, &gin.H{"token": userToken})
}
