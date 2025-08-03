package routes

import (
	"ayam-geprek-backend/controllers"
	"ayam-geprek-backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func DashboardRoutes(app *fiber.App) {
	// Protected dashboard group
	dashboard := app.Group("/dashboard", middlewares.JWTProtectedHTML())

	dashboard.Get("/", controllers.IndexDashboard)
}
