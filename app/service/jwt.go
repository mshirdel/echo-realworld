package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("secret-key")

type UserInfo struct {
	Username string
	UserID   uint
	ExpDate  int64
	jwt.Claims
}

func CreateToken(username string, id uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserInfo{
		Username: username,
		UserID:   id,
		ExpDate:  time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*UserInfo, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	if claims, ok := token.Claims.(UserInfo); ok {
		return &claims, nil
	}

	return nil, nil
}
