package user

import (
	"backend/myerr"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func Identity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		myerr.NewErrorResponse(c, http.StatusUnauthorized, "empty authorization header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		myerr.NewErrorResponse(c, http.StatusUnauthorized, "invalid authorization header")
		return
	}
	userId, err := ParseToken(headerParts[1])
	if err != nil {
		myerr.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	if !IsUserExist(userId) {
		myerr.NewErrorResponse(c, http.StatusForbidden, "user access denied")
		return
	}
	c.Set(userCtx, userId)
}

// currentUser()
// currentUserId()
// checkAccessGranted()

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		myerr.NewErrorResponse(c, http.StatusNotFound, "user id not found")
		return 0, errors.New("user id not found")
	}
	idInt, ok := id.(int)
	if !ok {
		myerr.NewErrorResponse(c, http.StatusNotFound, "user id is not int type")
		return 0, errors.New("user id is not int type")
	}
	return idInt, nil
}
