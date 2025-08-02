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

// ✅ Updated untuk return: token + user
func LoginUser(input LoginInput) (string, *models.User, error) {
	userPtr, err := models.GetUserByUsername(input.Username)
	if err != nil || userPtr == nil {
		return "", nil, errors.New("invalid credentials")
	}
	user := *userPtr

	if !models.CheckPassword(input.Password, user.Password) {
		return "", nil, errors.New("wrong password")
	}

	// Buat JWT claim
	claims := middlewares.JwtCustomClaims{
		Id_user:  user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	// Buat JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(middlewares.JwtSecret)
	if err != nil {
		return "", nil, err
	}

	// Jangan kirim password ke response
	user.Password = ""

	return signedToken, &user, nil
}
