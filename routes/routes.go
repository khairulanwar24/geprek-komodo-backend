package routes

import (
	"ayam-geprek-backend/controllers"
	"ayam-geprek-backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/", controllers.ShowLoginPage)
	app.Get("/dashboard", middlewares.AuthMiddleware, controllers.DashboardPage)

	SetupRoutesAuth(app)
	SetupRoutesStock(app)
	SetupRoutesOutlets(app)
	SetupRoutesKeuangan(app)
	SetupRoutesLaporan(app)

	// halaman login
	// routes
	// app.Get("/dashboard", middlewares.JWTProtected(), controllers.DashboardPage)

	// auth := app.Group("/api", middlewares.JWTProtected())

	// auth.Post("/stock", controllers.CreateStock)
	// auth.Get("/stock", controllers.ListStock)
}
