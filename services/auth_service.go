package services

import (
	"ayam-geprek-backend/middlewares"
	"ayam-geprek-backend/models"

	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginUser(input LoginInput) (string, error) {
	userPtr, err := models.GetUserByUsername(input.Username)
	if err != nil || userPtr == nil {
		return "", errors.New("invalid credentials")
	}
	user := *userPtr

	if !models.CheckPassword(input.Password, user.Password) {
		return "", errors.New("wrong password")
	}

	claims := middlewares.JwtCustomClaims{
		Id_user:  user.ID,
		Username: user.Username,

		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(middlewares.JwtSecret)

}
