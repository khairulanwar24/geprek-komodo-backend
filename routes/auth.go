package routes

import (
	"ayam-geprek-backend/controllers"
	"ayam-geprek-backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutesAuth(app *fiber.App) {

	app.Get("/", controllers.ShowLoginPage)
	// 🔓 PUBLIC: Tidak pakai JWT middleware
	app.Post("/auth/login", controllers.Login)
	app.Post("/auth/logout", controllers.Logout)

	// 🔐 PROTECTED: Group ini pakai JWT middleware
	auth := app.Group("/auth", middlewares.JWTProtected())
	auth.Get("/me", controllers.Me)
}
