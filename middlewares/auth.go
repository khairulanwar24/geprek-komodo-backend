package middlewares

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	accessToken := c.Cookies("access_token")

	if accessToken == "" {
		// Jika tidak ada token, redirect ke login
		return c.Redirect("/auth/login")
	}

	// Decode token (opsional, kalau kamu pakai JWT lokal)
	payload, err := DecodeJWT(accessToken)
	if err != nil {
		// Jika token tidak valid
		return c.Redirect("/auth/login")
	}

	// Cek expiry token (opsional kalau pakai JWT lokal)
	if exp, ok := payload["exp"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			return c.Redirect("/auth/login")
		}
	}

	return c.Next()
}

func DecodeJWT(token string) (map[string]interface{}, error) {
	parts := strings.Split(token, ".")
	if len(parts) < 2 {
		return nil, fmt.Errorf("token format salah")
	}

	payloadPart := parts[1]
	// tambahkan padding
	if m := len(payloadPart) % 4; m != 0 {
		payloadPart += strings.Repeat("=", 4-m)
	}

	payloadBytes, err := base64.URLEncoding.DecodeString(payloadPart)
	if err != nil {
		return nil, err
	}

	var payload map[string]interface{}
	if err := json.Unmarshal(payloadBytes, &payload); err != nil {
		return nil, err
	}

	return payload, nil
}
