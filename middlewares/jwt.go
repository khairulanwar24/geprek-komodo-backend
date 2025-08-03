package middlewares

import (
	"ayam-geprek-backend/types"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// 👇 Secret disimpan di sini
var JwtSecret = []byte("halosayang")

// 👇 Custom claims
type JwtCustomClaims struct {
	Username string    `json:"username"`
	Id_user  uuid.UUID `json:"id"`
	jwt.RegisteredClaims
}

// 🔑 Token Generator
func GenerateAccessToken(username string, id_user uuid.UUID) (string, error) {
	claims := &JwtCustomClaims{
		Username: username,
		Id_user:  id_user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var tokenStr string

		// Coba ambil dari header Authorization
		authHeader := c.Get("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
			tokenStr = strings.TrimPrefix(authHeader, "Bearer ")
		} else {
			// Coba ambil dari cookie
			tokenStr = c.Cookies("access_token")
		}

		// 🔍 DEBUG LOG
		// fmt.Println("Token from cookie/header:", tokenStr)

		if tokenStr == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(types.Response{
				Success: false,
				Message: "Missing or invalid token",
				Data:    nil,
			})
		}

		token, err := jwt.ParseWithClaims(tokenStr, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return JwtSecret, nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(types.Response{
				Success: false,
				Message: "Invalid or expired token",
				Data:    nil,
			})
		}
		claims := token.Claims.(*JwtCustomClaims)

		c.Locals("username", claims.Username)
		c.Locals("id_user", claims.Id_user)

		return c.Next()
	}
}

func JWTProtectedHTML() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var tokenStr string
		authHeader := c.Get("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
			tokenStr = strings.TrimPrefix(authHeader, "Bearer ")
		} else {
			tokenStr = c.Cookies("access_token")
		}

		if tokenStr == "" {
			return c.Redirect("/") // ke login
		}

		token, err := jwt.ParseWithClaims(tokenStr, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return JwtSecret, nil
		})

		if err != nil || !token.Valid {
			return c.Redirect("/")
		}
		claims := token.Claims.(*JwtCustomClaims)
		c.Locals("username", claims.Username)
		c.Locals("id_user", claims.Id_user)
		return c.Next()
	}
}
