package controllers

import (
	"ayam-geprek-backend/services"
	"ayam-geprek-backend/types"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Me(c *fiber.Ctx) error {
	username := c.Locals("username")
	idUser := c.Locals("id_user")

	return c.JSON(types.Response{
		Success: true,
		Message: "Authenticated user",
		Data: fiber.Map{
			"username": username,
			"id_user":  idUser,
		},
	})
}

func Login(c *fiber.Ctx) error {
	var input services.LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(types.Response{
			Success: false,
			Message: "Invalid input",
			Data:    nil,
		})
	}

	token, err := services.LoginUser(input)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(types.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// ✅ Set token as cookie
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    token,
		HTTPOnly: true,
		SameSite: "Lax",
		Path:     "/",
	})

	return c.JSON(types.Response{
		Success: true,
		Message: "Login successful",
		Data:    fiber.Map{"token": token},
	})
}

func Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour), // set ke waktu lampau
		HTTPOnly: true,
		SameSite: "Lax",
		Path:     "/",
	})
	return c.JSON(types.Response{
		Success: true,
		Message: "Logged out successfully",
		Data:    nil,
	})
}
