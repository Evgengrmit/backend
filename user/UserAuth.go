package user

import (
	"backend/db"
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

const (
	signedKey = "nx74nfiuyfgd4ju7j8ref74e"
	tokenTTL  = 12 * time.Hour
)

type TokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

func GenerateToken(data *LoginData) (string, error) {
	u, err := LoginUser(data)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix()},
		int(u.ID),
	})
	tokenStr, err := token.SignedString([]byte(signedKey))
	if err != nil {
		return "", err
	}
	_, err = db.DB.Exec("insert into token (user_id,token,is_valid) values ($1,$2,$3)", u.ID, tokenStr, true)
	if err != nil {
		return "", err
	}
	return token.SignedString([]byte(signedKey))
}

func ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signedKey), nil
	})
	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims.UserID, nil
	}
	return 0, errors.New("token claims are not of type *tokenClaims")
}
