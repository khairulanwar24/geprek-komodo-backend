package controllers

import (
	"ayam-geprek-backend/services"
	"ayam-geprek-backend/types"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ShowLoginPage(c *fiber.Ctx) error {
	return c.Render("login/login", fiber.Map{
		"Title": "Login",
	})
}

func Login(c *fiber.Ctx) error {
	// Tangkap input dari form HTML atau JSON API
	var input services.LoginInput
	if err := c.BodyParser(&input); err != nil {
		fmt.Println("[DEBUG] BodyParser error:", err)

		if c.Get("Content-Type") == "application/json" {
			return c.Status(fiber.StatusBadRequest).JSON(types.Response{
				Success: false,
				Message: "Invalid input",
			})
		}
		return c.Render("login/login", fiber.Map{
			"Error": "Input tidak valid",
		}, "layouts/base")
	}

	// Debug input yang diterima
	fmt.Println("[DEBUG] Login input - Username:", input.Username, "Password:", input.Password)

	token, user, err := services.LoginUser(input)
	if err != nil {
		fmt.Println("[DEBUG] LoginUser error:", err)

		if c.Get("Content-Type") == "application/json" {
			return c.Status(fiber.StatusUnauthorized).JSON(types.Response{
				Success: false,
				Message: err.Error(),
			})
		}
		return c.Render("login/login", fiber.Map{
			"Error": err.Error(),
		}, "layouts/base")
	}

	fmt.Println("[DEBUG] Login success - Token:", token)

	// Set cookie token
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    token,
		HTTPOnly: true,
		SameSite: "Lax",
		Path:     "/",
	})

	// ✅ Kalau dari API (React / JSON), kirim token + user
	if c.Get("Content-Type") == "application/json" {
		return c.JSON(types.Response{
			Success: true,
			Message: "Login successful",
			Data: fiber.Map{
				"token": token,
				"user": fiber.Map{
					"id_user":  user.ID,
					"nama":     user.Nama,
					"username": user.Username,
					"email":    user.Email,
					"no_hp":    user.NoHp,
					"role":     user.Role,
				},
			},
		})
	}

	// ✅ Kalau dari HTML, redirect ke dashboard
	return c.Redirect("/dashboard")
}

func Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HTTPOnly: true,
		SameSite: "Lax",
		Path:     "/",
	})

	// Jika HTML: redirect ke login
	if c.Get("Content-Type") != "application/json" {
		return c.Redirect("/login")
	}

	// Jika API: balikan JSON
	return c.JSON(types.Response{
		Success: true,
		Message: "Logged out successfully",
	})
}

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
