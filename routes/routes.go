package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	SetupRoutesAuth(app)
	SetupRoutesStock(app)

	// auth := app.Group("/api", middlewares.JWTProtected())

	// auth.Post("/stock", controllers.CreateStock)
	// auth.Get("/stock", controllers.ListStock)
}
