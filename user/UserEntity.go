package user

import (
	"backend/account"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID             uint64
	Name           string            `json:"name"`
	Age            int8              `json:"age,omitempty"`
	Login          string            `json:"login"`
	Email          string            `json:"email"`
	HashedPassword []byte            `json:"-"`
	Accounts       []account.Account `json:"accounts,omitempty"`
}

type Users struct {
	Users []User
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
